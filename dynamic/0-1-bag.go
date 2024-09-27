package dynamic

// 0-1 背包问题
// 物品0: weight=1, value=15
// 物品1: weight=2, value=20
// 物品2: weight=4, value=30

// dp[i][j]: 前i个物品，背包容量为j时的最大价值
func MaxValueForBag(weights []int, values []int, bagWeight int) int {
	if len(weights) != len(values) {
		return -1
	}

	dp := make([][]int, len(weights))
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}

	for j := weights[0]; j <= bagWeight; j++ {
		dp[0][j] = values[0]
	}

	for i := 1; i < len(weights); i++ { // 物品
		for j := 0; j <= bagWeight; j++ { // 背包
			if j < weights[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i]]+values[i])
			}
		}
	}

	return dp[len(weights)-1][bagWeight]
}

func MaxValueForBagII(weights []int, values []int, bagWeight int) int {
	if len(weights) != len(values) {
		return -1
	}

	dp := make([][]int, len(weights))
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}

	for j := weights[0]; j <= bagWeight; j++ {
		dp[0][j] = values[0]
	}
	for j := 0; j <= bagWeight; j++ { // 背包
		for i := 1; i < len(weights); i++ { // 物品
			if j < weights[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i]]+values[i])
			}
		}
	}

	return dp[len(weights)-1][bagWeight]
}

func MaxValueForBagIII(weights []int, values []int, bagWeight int) int {
	if len(weights) != len(values) {
		return -1
	}

	dp := make([]int, bagWeight+1)

	for i := 0; i < len(weights); i++ { // 物品
		for j := bagWeight; j >= weights[i]; j-- { //背包
			dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
		}
	}

	return dp[bagWeight]
}
