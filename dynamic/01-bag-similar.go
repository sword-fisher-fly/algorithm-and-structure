package dynamic

import "math"

// 划归为01背包问题
// 1）石头相毁, 求最后剩下的一块石头的重量
// 2）目标和, 求一组数字前面放置'+'和'-'号使之和等于目标有多少种方法
// 3）0和1, 求一组字符串都是由0和1数字构成,最多包含m个0和n个1的最大子集大小

// 完全背包问题,
// I: 先背包, 后物品 II: 先物品, 后背包
// 1) 单词拆分
// 2) 零钱兑换
// 如果求组合数就是外层for循环遍历物品，内层for遍历背包
// 如果求排列数就是外层for遍历背包，内层for循环遍历物品

// 目标和: 求所有添加符号的方法数

// dp[j]: 总和为j的方法数
// left + right = sum
// left - right = target
// bagSize = (target+sum)/2
func FindTargetSumWays(nums []int, target int) int {
	abs := func(a int) int {
		if a > 0 {
			return a
		}

		return -a
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if abs(target) > sum {
		return 0
	}

	if (sum+target)%2 == 1 {
		return 0
	}

	bagSize := (sum + target) / 2

	dp := make([]int, bagSize+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ { // 物品
		for j := bagSize; j >= nums[i]; j-- { // 背包
			// for j := nums[i]; j <= bagSize; j++ { // NOT PASS
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[bagSize]
}

func FindTargetSumWaysByBacktracing(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if target > sum {
		return 0
	}

	if (sum+target)%2 == 1 {
		return 0
	}

	// var res [][]int

	var res int
	bagSize := (sum + target) / 2

	// fmt.Printf("sum=%d, target=%d, bagSize=%d\n", sum, target, bagSize)

	var backtracing func(candidates []int, startIndex int, path []int, sum int)
	backtracing = func(candidates []int, startIndex int, path []int, sum int) {
		if sum == bagSize {
			// tmp := make([]int, len(path))
			// copy(tmp, path)
			// res = append(res, tmp)
			res++
		}

		// fmt.Printf("in backtracing, sum=%d\n", sum)

		for i := startIndex; i < len(candidates) && sum+candidates[i] <= bagSize; i++ {
			sum += candidates[i]
			// path = append(path, candidates[i])
			backtracing(candidates, i+1, path, sum)
			sum -= candidates[i]
			// path = path[:len(path)-1]
		}
	}

	backtracing(nums, 0, []int{}, 0)

	// return len(res)
	return res
}

// dp[j]: 兑换总额为j的总方法数
func CoinChange(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ { // 物品
		for j := coins[i]; j <= amount; j++ { // 背包
			// for j := amount; j >= coins[i]; j-- { // NOT PASS
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}

// dp[i]: 兑换总额为i的最小硬币数， coins数量都是1个？？
func CoinChangeWithMinCount(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = 0
	for i := 0; i < len(coins); i++ { //物品
		for j := coins[i]; j <= amount; j++ { //背包
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}

	if dp[amount] == math.MaxInt {
		return 0
	}

	return dp[amount]
}

func CoinChangeWithMinCountII(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ { // 背包
		for j := 0; j < len(coins); j++ { // 物品
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return 0
	}

	return dp[amount]
}
