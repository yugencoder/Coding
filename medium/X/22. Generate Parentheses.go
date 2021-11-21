package X

func generateParenthesis(n int) []string {
	res := []string{}
	generate(0, n, 0, 0, "", &res)

	return res
}

func generate(curr, end, l, r int, currVal string, res *[]string) {
	if curr == 2*end {
		*res = append(*res, currVal)
		return
	}

	opts := []string{"(", ")"}

	for _, o := range opts {
		if o == "(" && l < end {
			generate(curr+1, end, l+1, r, currVal+o, res)
		} else if o == ")" && r < end && r < l {
			generate(curr+1, end, l, r+1, currVal+o, res)
		}
	}
}
