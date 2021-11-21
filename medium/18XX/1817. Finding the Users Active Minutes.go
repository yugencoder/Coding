package _8XX

import "sort"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	res := make([]int, k)
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	prev := logs[0][0]
	uam := 0
	uamMap := map[int]bool{}
	for i, l := range logs {
		if l[0] != prev && uam > 0{
			res[uam-1]++
			uamMap = map[int]bool{}
			uam = 0
		}

		if !uamMap[l[1]] {
			uam += 1
		}

		uamMap[l[1]] = true
		prev = l[0]

		if i == len(logs)-1 && uam > 0 {
			res[uam-1]++
		}
	}

	return res
}
