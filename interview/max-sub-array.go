package interview

// TODO: https://leetcode.cn/problems/maximum-subarray/
// 1) dp
func MaxSubArrayByDP(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	pre, maxSum := 0, nums[0]
	for i := range nums {
		pre = max(pre+nums[i], nums[i])
		maxSum = max(maxSum, pre)
	}

	return maxSum
}

// 2) greedy

func MaxSubArrayByGreedy(nums []int) int {
	sum, maxSum := 0, nums[0]

	for i := range nums {
		sum += nums[i]

		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}

// 3) prefix sum
func MaxSubArrayByPrefixSum(nums []int) int {
	maxSum := nums[0]
	preSum, minPreSum := 0, 0

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	for i := range nums {
		preSum += nums[i]
		maxSum = max(maxSum, preSum-minPreSum)
		minPreSum = min(minPreSum, preSum)
	}

	return maxSum
}
