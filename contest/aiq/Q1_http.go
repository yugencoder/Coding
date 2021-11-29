package aiq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func main() {
	articles := getAllArticles()

	fmt.Printf("Got %d articles\n", len(articles))

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].NumComments <= articles[i].NumComments
	})

	for _, a := range articles {
		fmt.Println(a)
	}

}

func getAllArticles() []Article {
	totalPages := 1
	articles := []Article{}

	for i := 0; i < totalPages; i++ {
		articleResponse := getArticleResponse(i)
		articles = append(articles, articleResponse.Data...)

		if i == 0 {
			totalPages = articleResponse.TotalPages
		}
	}
	return articles
}

func getArticleResponse(pageNo int) (articleResponse ArticleResponse) {
	articleURL := strings.Join([]string{articleHost, articlePath}, "/")

	req, err := http.NewRequest("GET", articleURL, nil)
	if err != nil {
		fmt.Printf("[ERROR] Cannot get articles, %s returned %d  \n", articleURL, err)
		return
	}

	q := req.URL.Query()
	q.Add("page", strconv.Itoa(pageNo))
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[ERROR] Cannot make articles request, client returned : %s", err)
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&articleResponse); err != nil {
		fmt.Printf("[ERROR] Cannot decode %s response : %s", articleURL, err)
		return
	}

	return articleResponse
}

type ArticleResponse struct {
	Page       int       `json:"page"`
	PerPage    int       `json:"per_page"`
	Total      int       `json:"total"`
	TotalPages int       `json:"total_pages"`
	Data       []Article `json:"data"`
}

type Article struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Author      string `json:"author"`
	NumComments int    `json:"num_comments"`
	StoryID     int32  `json:"story_id"`
	StoryTitle  string `json:"story_title"`
	StoryURL    string `json:"story_url"`
	ParentID    int32  `json:"parent_id"`
	CreatedAt   int    `json:"created_at"`
}

type CommentTitle struct {
	NumberOfComments int    `json:"number_of_comments"`
	Title            string `json:"title"`
	Page             int    `json:"from_page"`
}
