package aiq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)



const (
	articleHost = "https://jsonmock.hackerrank.com"
	articlePath = "api/articles"
)

var wg sync.WaitGroup

func main2() {
	comment_title_chan := make(chan CommentTitle)
	var commentTitleSlice []CommentTitle

	// pilot call to get total number of pages
	wg.Add(1)
	totalPage := makePaginatedRequest(1, comment_title_chan, true)

	// making concurrent requests to multiple pages at once
	for j := 1; j <= totalPage; j++ {
		wg.Add(1)
		go makePaginatedRequest(j, comment_title_chan, false)
	}

	for {
		commentTitleSlice = append(commentTitleSlice, <-comment_title_chan)
	}
	go func() {
		wg.Wait()
		close(comment_title_chan)
	}()

	for _, article := range commentTitleSlice {
		fmt.Println(article.NumberOfComments, "\t\t", article.Title)
	}
}

func makePaginatedRequest(pageNo int, chunk chan CommentTitle, pilotMode bool) int {
	defer wg.Done()
	uri := articleHost + articlePath
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println(err)
	}
	q := req.URL.Query()
	q.Add("page", strconv.Itoa(pageNo))
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	var articleResponse ArticleResponse
	if err = json.NewDecoder(resp.Body).Decode(&articleResponse); err != nil {
		fmt.Println(err)
	}
	if !pilotMode {
		for _, article := range articleResponse.Data {
			if article.Title != "" && article.NumComments != 0 {
				ct := CommentTitle{article.NumComments, article.Title, pageNo}
				chunk <- ct
			}
		}
	}
	return articleResponse.TotalPages
}
