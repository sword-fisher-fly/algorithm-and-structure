package offer

// TODO:
// https://leetcode.cn/problems/target-sum/solutions/816361/mu-biao-he-by-leetcode-solution-o0cp/

func FindTargetSum(nums []int, target int) int {
	var backtracing func(idx, sum int)
	count := 0
	backtracing = func(idx, sum int) {
		if idx == len(nums) {
			if sum == target {
				count++
			}
			return
		}

		backtracing(idx+1, sum+nums[idx])
		backtracing(idx+1, sum-nums[idx])
	}

	backtracing(0, 0)

	return count
}

/*
因此状态转移方程如下：
dp[i][j]=dp[i-1][j], if j < nums[i]
dp[i][j]=dp[i−1][j]+dp[i−1][j−nums[i]], if j ≥ nums[i]
*/

// (sum-neg)-neg = target -> neg = (sum-target)>>1
func FindTargetSumByDP(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	diff := sum - target
	if diff < 0 || diff&1 != 0 {
		return 0
	}

	n, neg := len(nums), diff>>1
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}

	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		num := nums[i-1]
		for j := 0; j <= neg; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= num {
				dp[i][j] += dp[i-1][j-num]
			}
		}
	}

	return dp[n][neg]
}
