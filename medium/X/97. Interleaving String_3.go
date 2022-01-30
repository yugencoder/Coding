package X

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] += 1
			} else if i == 0 && s2[j-1] == s3[i+j-1] && dp[i][j-1] > 0 {
				dp[i][j] += 1
			} else if j == 0 && s1[i-1] == s3[i+j-1] && dp[i-1][j] > 0 {
				dp[i][j] += 1
			} else if i > 0 && s1[i-1] == s3[i+j-1] && dp[i-1][j] > 0 || j > 0 && s2[j-1] == s3[i+j-1] && dp[i][j-1] > 0 {
				dp[i][j] += 1
			}
		}
	}

	return dp[len(s1)][len(s2)] > 0
}
