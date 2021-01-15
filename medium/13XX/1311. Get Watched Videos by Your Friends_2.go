package t13XX

import sort "sort"

func watchedVideosByFriends2(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	already := make(map[int]bool)
	already[id] = true

	queue := make([]int, 1)
	queue[0] = id
	// This can be set to any value(as long as it is not too large)
	initCap := len(friends) / level

	for i := 0; i < level && len(queue) != 0; i++ {
		nextQueue := make([]int, 0, initCap)
		for _, f := range queue {
			for _, ff := range friends[f] {
				if _, ok := already[ff]; !ok {
					already[ff] = true
					nextQueue = append(nextQueue, ff)
				}
			}
		}
		queue = nextQueue
	}

	videoCounter := make(map[string]int)
	for _, f := range queue {
		for _, v := range watchedVideos[f] {
			if c, ok := videoCounter[v]; !ok {
				videoCounter[v] = 1
			} else {
				videoCounter[v] = c + 1
			}
		}
	}

	pairList := make(PairList, 0, len(videoCounter))
	for k, v := range videoCounter {
		pairList = append(pairList, Pair{v, k})
	}
	sort.Sort(pairList)

	res := make([]string, len(pairList))
	for i, p := range pairList {
		res[i] = p.Video
	}
	return res
}

type Pair struct {
	Count int
	Video string
}

type PairList []Pair

func (list PairList) Len() int {
	return len(list)
}

func (list PairList) Less(i, j int) bool {
	if list[i].Count == list[j].Count {
		return list[i].Video < list[j].Video
	} else {
		return list[i].Count < list[j].Count
	}
}

func (list PairList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}