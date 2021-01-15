package _14XX

func getHappyString(n int, k int) string {
	strs := getAllHappyStrings(0, n, []string{"a", "b", "c"})

	if k > len(strs){
		return ""
	}

	return strs[k-1]
}

func getAllHappyStrings(i int, n int, res []string) []string {
	next := []string{"a", "b", "c"}

	if i == n-1 {
		return res
	}

	newRes := []string{}
	for _, r := range res {
		for _, n := range next {
			if r[len(r)-1] != n[0] {
				newRes = append(newRes, r+n)
			}
		}
	}

	return getAllHappyStrings(i+1, n, newRes)
}
