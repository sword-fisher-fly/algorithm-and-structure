package dynamic

// https://leetcode.cn/problems/interleaving-string/
func IsInterleave(s1 string, s2 string, s3 string) bool {
	m1, m2, m3 := len(s1), len(s2), len(s3)
	if m1+m2 != m3 {
		return false
	}

	dp := make([][]bool, m1+1)
	for i := range dp {
		dp[i] = make([]bool, m2+1)
	}

	dp[0][0] = true
	for i := 0; i < m1+1; i++ {
		for j := 0; j < m2+1; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}

			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}

	return dp[m1][m2]
}
