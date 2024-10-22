package offer

// TODO:
// https://leetcode.cn/problems/counting-bits/

func BitCount(n int) int {
	count := 0
	for n != 0 {
		count++
		n = n & (n - 1)
	}

	return count
}

func CountBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = BitCount(i)
	}

	return ans
}

func CountBitsII(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + i&1
	}

	return ans[1:]
}

func CountBitsIII(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i&(i-1)] + 1
	}

	return ans[1:]
}
