package greedy

// 1) 遍历计算0..i的和, 遇到sum<0, 重设为0开始计算
func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	sum := 0
	for _, num := range nums {
		sum += num
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}

// 2）动态规划
// dp[i]: 以nums[i]结尾的最大子数组和
func MaxSubArrayII(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}
