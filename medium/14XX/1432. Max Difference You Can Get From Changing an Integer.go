package _14XX

import (
	"fmt"
	"strconv"
	"strings"
)

func maxDiff(num int) int {
	mmap := map[int]bool{}
	for t := num; t > 0; t /= 10 {
		d := t % 10
		mmap[d] = true
	}

	numS := strconv.Itoa(num)
	if len(mmap) == 1 {
		numS = strings.Repeat("8", len(numS))
		n, _ := strconv.Atoi(numS)
		return n
	}
	max := getMaxNo(num)
	min := getMinNo(num)

	fmt.Println(max)
	fmt.Println(min)
	return max - min
}

func getMinNo(num int) int {
	numS := strconv.Itoa(num)

	for i, c := range numS {
		c := string(c)
		if i == 0 && c != "1" {
			nNumS := strings.Replace(numS, string(numS[i]), "1", -1)
			n, _ := strconv.Atoi(nNumS)
			return n
		} else if i > 0 {
			if c != "0" && c != string(numS[0]) {
				nNumS := strings.Replace(numS, string(numS[i]), "0", -1)
				n, _ := strconv.Atoi(nNumS)
				return n
			}
		}
	}

	return num
}

func getMaxNo(num int) int {
	numS := strconv.Itoa(num)
	for i, c := range numS {
		c := string(c)
		if c != "9" {
			nNumS := strings.Replace(numS, string(numS[i]), "9", -1)
			n, _ := strconv.Atoi(nNumS)
			return n
		}
	}
	return num
}

func reverseSlice(digits []int) {
	for i, j := 0, len(digits)-1; i < len(digits)/2; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
}
