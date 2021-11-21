package X

//func isInterleave(s1 string, s2 string, s3 string) bool {
//	if len(s1)+len(s2) != len(s3) {
//		return false
//	}
//	return isValidInterleave(0, 0, 0, s1, s2, s3, "")
//}
//
//func isValidInterleave(i, j, k int, s1 string, s2 string, s3 string, s4 string) bool {
//	if k == len(s3) && i == len(s1) && j == len(s2) {
//		return true
//	}
//
//	if k >= len(s3) {
//		return false
//	}
//
//	if i < len(s1) && s1[i] == s3[k] {
//		v1 := isValidInterleave(i+1, j, k+1, s1, s2, s3, s4)
//		if v1 {
//			return true
//		}
//	}
//
//	if j < len(s2) && s2[j] == s3[k] {
//		v1 := isValidInterleave(i, j+1, k+1, s1, s2, s3, s4)
//		if v1 {
//			return true
//		}
//	}
//
//	return false
//
//}
