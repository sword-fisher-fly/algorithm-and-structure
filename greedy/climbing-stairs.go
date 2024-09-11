package greedy

// dp[i]: 爬上第i层所需的最小花费
func MinCostClimbingStairs(cost []int) int {
	if len(cost) < 3 {
		return 0
	}

	dp := make([]int, len(cost)+1)

	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}

func MinCostClimbingStairsOptimize(cost []int) int {
	if len(cost) < 3 {
		return 0
	}

	dp0, dp1 := 0, 0
	for i := 2; i <= len(cost); i++ {
		dpi := min(dp1+cost[i-1], dp0+cost[i-2])
		dp0 = dp1
		dp1 = dpi
	}

	return dp1
}
