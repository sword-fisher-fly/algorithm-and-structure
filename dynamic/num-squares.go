package dynamic

import "math"

// https://leetcode.cn/problems/perfect-squares/?envType=problem-list-v2&envId=2cktkvj
func NumSquares(n int) int {
	dp := make([]int, n+1)

	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	for i := 1; i <= n; i++ {
		minn := math.MaxInt

		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}

		dp[i] = minn + 1 // i - j*j
	}

	return dp[n]
}
