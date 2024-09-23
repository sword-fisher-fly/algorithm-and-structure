package dynamic

// 假设nums的所有元素都为正整数
func CanPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum&1 == 1 {
		return false
	}

	target := sum >> 1

	dp := make([]int, target+1)

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[target] == target
}
