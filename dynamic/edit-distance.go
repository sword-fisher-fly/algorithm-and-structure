package dynamic

// word1, word2 找到使得word1和word2相同所需的最小步数
// 每步可以删除任意一个字符串的一个字符

// max(i): len(word1)-1
// max(j): len(word2)-1
// dp[i][j]: word1[:i+1] 和 word2[:j+1] 的最小编辑距离
func MinDeleteDistance(word1, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	dp := make([][]int, len(word1))
	for i := range dp {
		dp[i] = make([]int, len(word2))
	}

	for i := range dp {
		dp[i][0] = i
	}

	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[len(word1)-1][len(word2)-1]
}

// 化归为最长公共子序列
// 最小步数 = len(word1) + len(word2) - 2*LCS(word1,word2)
// dp[i][j]: 表示 word1[:i] 和 word2[:j] 的最长公共子序列的长度
func MinDeleteDistanceII(word1, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return len(word1) + len(word2) - 2*dp[len(word1)][len(word2)]
}

// word1, word2 找到使得word1和word2相同所需的最小步数
// 每一步只能新增、删除、替换
// dp[i][j]: word1[:i] 和 word2[:j] 的最小编辑距离
func MinEditDistance(word1, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}

	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
