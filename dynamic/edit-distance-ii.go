package dynamic

// 1.插入一个字符
// 2.删除一个字符
// 3.修改一个字符
func MinEditDistanceWithInsertDeleteReplace(str1 string, str2 string) int {
	l1, l2 := len(str1), len(str2)

	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := range dp {
		dp[i][0] = i
	}

	for i := range dp[0] {
		dp[0][i] = i
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[l1][l2]
}
