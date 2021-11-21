package _14XX

import "fmt"

func simplifiedFractions(n int) []string {
	fractionMap := map[string]bool{}
	var res []string
	keyFormat := "%d/%d"
	for i := 1; i <= n; i++ {
		for d := i + 1; d <= n; d++ {

			if g := gcd(i, d); g > 1 {
				if i/g == d/g {
					continue
				}

				fractionMap[fmt.Sprintf(keyFormat, i/g, d/g)] = true
				continue
			}

			fractionMap[fmt.Sprintf(keyFormat, i, d)] = true
		}
	}

	for k, _ := range fractionMap {
		res = append(res, k)
	}

	return res
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
