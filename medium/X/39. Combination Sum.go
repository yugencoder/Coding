package X

var res2 [][]int

func combinationSum(candidates []int, target int) [][]int {
	res2 = [][]int{}

	for i, c := range candidates {
		findTarget([]int{c}, i, candidates, target-c)
	}

	return res2
}

func findTarget(currRes []int, i int, candidates []int, target int) {
	if target == 0 {
		newRes := make([]int, len(currRes))
		copy(newRes, currRes)
		res2 = append(res2, newRes)
		return
	} else if target < 0 {
		return
	}

	for j := i; j < len(candidates); j++ {
		findTarget(append(currRes, candidates[j]), j, candidates, target-candidates[j])
	}
}
