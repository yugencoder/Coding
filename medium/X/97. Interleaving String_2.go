package X

//func isInterleave(s1 string, s2 string, s3 string) bool {
//	return isValidInterleave(0, 0, 0, s1, s2, s3)
//}
//
//func isValidInterleave(i, j, k int, s1 string, s2 string, s3 string) bool {
//	if k == len(s3) && i == len(s1) && j == len(s2) {
//		return true
//	}
//
//	i1 := i
//	if k1 := k; i1 < len(s1) && k1 < len(s3) && s1[i1] == s3[k1] {
//		for i1 < len(s1) && k1 < len(s3) && s1[i1] == s3[k1] {
//			isValidS1_ := isValidInterleave(i1+1, j, k1+1, s1, s2, s3)
//			if isValidS1_ {
//				return true
//			}
//			i1++
//			k1++
//		}
//	}
//
//	j1 := j
//	if k2 := k; j1 < len(s2) && k2 < len(s3) && s2[j1] == s3[k2] {
//		for j1 < len(s2) && k2 < len(s3) && s2[j1] == s3[k2] {
//			isValidS2_ := isValidInterleave(i, j1+1, k2+1, s1, s2, s3)
//			if isValidS2_ {
//				return true
//			}
//
//			j1++
//			k2++
//		}
//	}
//
//	//if i < len(s1) && k < len(s3) && s1[i] == s3[k] {
//	//	isValidS1 = isValidInterleave(i+1, j, k+1, s1, s2, s3)
//	//	if isValidS1 {
//	//		return true
//	//	}
//	//}
//	//
//	//if j < len(s2) && k < len(s3) && s2[j] == s3[k] {
//	//	return isValidInterleave(i, j+1, k+1, s1, s2, s3)
//	//}
//
//	return false
//
//}
