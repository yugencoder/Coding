package t13XX

func minSteps(s string, t string) int {
	charMap := map[rune]int{}
	res := len(t)

	for _,c := range s{
		charMap[c] += 1
	}


	for _,c:= range t{
		if v,ok:= charMap[c]; ok && v > 0{
			charMap[c] -= 1
			if v == 1{
				delete(charMap,c)
			}
			res--
		}
	}


	return res
}