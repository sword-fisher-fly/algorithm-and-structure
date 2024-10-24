package offer

// https://leetcode.cn/problems/min-cost-climbing-stairs/
// dp[i]: 到达下标 i 的最小花费
// dp[i]=min(dp[i−1]+cost[i−1],dp[i−2]+cost[i−2])
func MinCostClimbStairs(cost []int) int {
	if len(cost) <= 2 {
		return 0
	}

	dp := make([]int, len(cost)+1)

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}
