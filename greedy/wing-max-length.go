package greedy

// 最大摆动序列长度

// 1) 中间有相同值, 平坦的波峰或波谷
// 2) 中间有相同值, 单调非递增或非递减
// 3) 数组首尾两端
func WiggleMaxLength(nums []int) int {
	preDiff := 0
	curDiff := 0

	result := 1
	for i := 1; i < len(nums); i++ {
		curDiff = nums[i] - nums[i-1]
		if (curDiff > 0 && preDiff <= 0) || (curDiff < 0 && preDiff >= 0) {
			result++
			preDiff = curDiff
		}
	}

	return result
}

// dp[i][0]: 以nums[i-1]为结尾的谷底数量
// dp[i][1]: 以nums[i-1]为结尾的波峰数量
func WiggleMaxLengthByDP(nums []int) int {
	dp := make([][2]int, len(nums))

	dp[0][0], dp[0][1] = 1, 1
	for i := 1; i < len(nums); i++ {
		dp[i][0], dp[i][1] = 1, 1
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] { // 可能的波峰
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}
		}

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] { // 可能的谷底
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			}
		}
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}
