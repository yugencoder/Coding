package X

import "fmt"

var someMap map[string]map[string]bool

func isInterleave(s1 string, s2 string, s3 string) bool {
	if someMap == nil {
		someMap = map[string]map[string]bool{
			s1: {s2: false},
		}
	} else if someMap[s1] == nil {
		someMap[s1] = map[string]bool{s2: false}
	} else if someMap[s1] != nil {
		if v, ok := someMap[s1][s2]; ok {
			return v
		}
	}

	if len(s1)+len(s2) != len(s3) {
		return false
	}

	i := 0
	j := 0
	k := 0
	if i < len(s1) && j < len(s2) && s1[i] == s3[k] && s2[j] == s3[k] {
		someMap[s1[i+1:]][s2] = isInterleave(s1[i+1:], s2, s3[k+1:])
		fmt.Println(s1[i+1:])
		fmt.Println(s2)
		fmt.Println(s3[k+1:])
		someMap[s1][s2[j+1:]] = isInterleave(s1, s2[j+1:], s3[k+1:])
		fmt.Println(s1)
		fmt.Println(s2[j+1:])
		fmt.Println(s3[k+1:])
		return someMap[s1[i+1:]][s2] || someMap[s1][s2[j+1:]]
	} else if i < len(s1) && s1[i] == s3[k] {
		someMap[s1[i+1:]][s2] = isInterleave(s1[i+1:], s2, s3[k+1:])
		fmt.Println(s1[i+1:])
		fmt.Println(s2)
		fmt.Println(s3[k+1:])
		return someMap[s1[i+1:]][s2]
	} else if j < len(s2) && s2[j] == s3[k] {
		someMap[s1][s2[j+1:]] = isInterleave(s1, s2[j+1:], s3[k+1:])
		fmt.Println(s1)
		fmt.Println(s1, s2[j+1:])
		fmt.Println(s3[k+1:])
		return someMap[s1][s2[j+1:]]
	}

	return false
}
