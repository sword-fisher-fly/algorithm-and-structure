package interview

// dp[i][j]: 代表s[i..j]是否为回文字符串
func LongestPalindromeByDP(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	maxLen := 1
	start, end := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start, end = i, j
			}
		}
	}

	return s[start : end+1]
}

func LongestPalindromeByExpand(s string) string {
	maxLen := 1
	start, end := 0, 0

	var expand func(left, right int)

	expand = func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
				start, end = left, right
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}

	return s[start : end+1]
}

// TODO:
// 最小回文字符串切割数
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。

// 返回符合要求的 最少分割次数
// 分而治之
// 1) 回文标记
// 2）类似最大子序列求最小分割数
func MinPalindromeCut(s string) int {
	return 0
}
