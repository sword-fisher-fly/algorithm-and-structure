package dynamic

// 子数组连续
// dp[i][j]: 表示以下标 i-1 为结尾的 arr1 和以下标 j-1 为结尾的 arr2 的最长公共子数组的长度
func MaxCommonSubArray(arr1 []int, arr2 []int) int {
	dp := make([][]int, len(arr1)+1)
	for i := 0; i < len(arr1)+1; i++ {
		dp[i] = make([]int, len(arr2)+1)
	}

	result := 0
	for i := 1; i <= len(arr1); i++ {
		for j := 1; j <= len(arr2); j++ {
			if arr1[i-1] == arr2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			result = max(result, dp[i][j])
		}
	}

	return result
}

// 优化为一维
func MaxCommonSubArrayII(arr1 []int, arr2 []int) int {
	dp := make([]int, len(arr2)+1)

	result := 0
	for i := 1; i <= len(arr1); i++ {
		for j := len(arr2); j >= 1; j-- {
			if arr1[i-1] == arr2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0
			}

			result = max(result, dp[j])
		}
	}

	return result
}
