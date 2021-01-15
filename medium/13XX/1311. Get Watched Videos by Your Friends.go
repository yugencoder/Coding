package t13XX

import (
	"fmt"
	"sort"
)

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	var list []int
	watchedVideosMap := map[string]int{}

	traverseDFS(id, friends, 0, level, map[int]bool{}, &list)
	fmt.Println(list)
	//list = traverseBFS(id, friends, level)

	for _, l := range list {
		if l == id {
			continue
		}

		visitedMap := map[string]bool{}
		for _, wv := range watchedVideos[l] {
			if !visitedMap[wv] {
				watchedVideosMap[wv] += 1
			}
			visitedMap[wv] = true
		}
	}

	type watchedVideosStruct struct {
		watchedVideos string
		count         int
	}

	var wvs []watchedVideosStruct
	for k, v := range watchedVideosMap {
		wvs = append(wvs, watchedVideosStruct{
			watchedVideos: k,
			count:         v,
		})
	}

	sort.Slice(wvs, func(i, j int) bool {
		if wvs[i].count == wvs[j].count {
			return wvs[i].watchedVideos < wvs[j].watchedVideos
		}

		return wvs[i].count < wvs[j].count
	})

	var res []string
	for _, wv := range wvs {
		res = append(res, wv.watchedVideos)
	}

	return res
}

func traverseDFS(id int, friends [][]int, currLevel, level int, visitedNodes map[int]bool, arr *[]int) {
	visitedNodes[id] = true
	if currLevel == level {
		*arr = append(*arr, id)
		return
	}

	for _, l := range friends[id] {
		if !visitedNodes[l] {
			traverseDFS(l, friends, currLevel+1, level, visitedNodes, arr)
		}
	}

	return
}

func traverseBFS(id int, list [][]int, level int) []int {
	currentFriends := []int{id}
	visited := make([]bool, 100)

	for i := 1; i <= level; i++ {
		if len(list) == 0 {
			return []int{}
		}
		nextFriends := []int{}

		for _, fid := range currentFriends {
			visited[fid] = true
			for _, ffid := range list[fid] {
				if !visited[ffid] {
					visited[ffid] = true
					nextFriends = append(nextFriends, ffid)
				}
			}
		}
		currentFriends = nextFriends
	}

	return currentFriends
}
