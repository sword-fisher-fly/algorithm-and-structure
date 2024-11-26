package dynamic

// dp[i]以num[i]结尾的最长子序列长度
func LongestAscendingSubSequence(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}

	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = 1
	}

	result := 1

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		result = max(result, dp[i])
	}

	return result
}

// 2） 最长连续递增子序列
// dp[i]: 以nums[i]结尾的最长连续递增子序列
func LongestContinuousAscendingSubSequence(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	dp := make([]int, len(nums))

	for i := range dp {
		dp[i] = 1
	}

	result := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}

		result = max(result, dp[i])
	}

	return result
}
