package dynamic

func CountPalindromicSubsequence(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	result := 0
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- { // 从下往上
		for j := i; j < len(s); j++ { // 从左往右
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				result++
			}
		}
	}

	return result
}

func extendPalindromicSubsequence(s string, i, j int) int {
	result := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		result++
		i--
		j++
	}

	return result
}

func CountPalindromicSubsequenceII(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result += extendPalindromicSubsequence(s, i, i)
		result += extendPalindromicSubsequence(s, i, i+1)
	}

	return result
}
