package _14XX

func peopleIndexes(favoriteCompanies [][]string) []int {
	companyMap := map[string][]int{}

	for i, companies := range favoriteCompanies {
		for _, c := range companies {
			companyMap[c] = append(companyMap[c], i)
		}
	}

	res := []int{}
	for i, companies := range favoriteCompanies {
		temp := []int{}
		for _, c := range companies {
			if len(companyMap[c]) == 1 && companyMap[c][0] == i {
				res = append(res, i)
				break
			} else if len(companyMap[c]) > 1 {
				vals := intersection(temp, companyMap[c], i)
				if len(vals) == 0 {
					res = append(res, i)
					break
				} else {
					temp = vals
				}
			}
		}
	}

	return res
}

func intersection(a []int, b []int, curr int) []int {
	if len(a) == 0 {
		return b
	}

	var res []int
	for _, av := range a {
		if av == curr {
			continue
		}
		for _, bv := range b {
			if av == bv {
				res = append(res, av)
			}
		}
	}
	return res
}
