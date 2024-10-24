package offer

// TODO:
// https://leetcode.cn/problems/length-of-longest-fibonacci-subsequence/description/

func LenLongestFibSubseq(arr []int) int {
	n := len(arr)
	indices := make(map[int]int, n)
	for i, x := range arr {
		indices[x] = i
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// arr[j] * 2 > x 是个规律
	// 1,1,2,3,5,8
	ans := 0
	for i, x := range arr {
		for j := i - 1; j >= 0 && arr[j]*2 > x; j-- {
			if k, ok := indices[x-arr[j]]; ok {
				dp[j][i] = max(dp[k][j]+1, 3)
				ans = max(ans, dp[j][i])
			}
		}
	}

	return ans
}
