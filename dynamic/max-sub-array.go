package dynamic

func MaxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	maxVal := nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])

		if dp[i] > maxVal {
			maxVal = dp[i]
		}
	}

	return maxVal
}
