package dynamic

func MinMoney(arr []int, aim int) int {
	// write code here
	if aim < 1 {
		return 0
	}
	// 需要凑出i元需要的最小货币数量
	dp := make([]int, aim+1)
	for i := range dp {
		dp[i] = aim + 1
	}

	dp[0] = 0

	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}
	for i := 1; i <= aim; i++ { //背包
		for j := 0; j < len(arr); j++ { //物品
			if arr[j] <= i {
				dp[i] = min(dp[i], dp[i-arr[j]]+1)
			}
		}
	}

	if dp[aim] > aim {
		return -1
	}

	return dp[aim]
}

func helper(arr []int, aim int, dp []int) int {
	MIN_VALUE := aim + 1
	if aim < 0 {
		return -1
	}

	if aim == 0 {
		return 0
	}

	if dp[aim-1] != 0 {
		return dp[aim-1]
	}

	for i := 0; i < len(arr); i++ {
		res := helper(arr, aim-arr[i], dp)
		if res >= 0 && res < MIN_VALUE {
			MIN_VALUE = res + 1
		}
	}

	if MIN_VALUE == aim+1 {
		dp[aim-1] = -1
	} else {
		dp[aim-1] = MIN_VALUE
	}

	return dp[aim-1]
}

func MinMoneyByRecursive(arr []int, aim int) int {
	if aim < 1 {
		return 0
	}

	dp := make([]int, aim+1)

	return helper(arr, aim, dp)
}
