package dynamic

// dp[i][j]: 表示s[i:j+1]为回文字符串的最大长度
// 1) if s[i] == s[j], then dp[i][j] == dp[i+1][j-1]+2
// 2) if s[i] != s[j], then dp[i][j] = max(dp[i][j-1], dp[i+1][j])
func MaxSubPalindrome(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}

	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
	}

	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][len(s)-1]
}
