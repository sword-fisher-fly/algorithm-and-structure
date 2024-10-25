package offer

//https://leetcode.cn/problems/interleaving-string/description/

func IsInterleave(s1 string, s2 string, s3 string) bool {
	m1, m2, m3 := len(s1), len(s2), len(s3)
	if m1+m2 != m3 {
		return false
	}

	dp := make([][]int, m1+1)
	for i := range dp {
		dp[i] = make([]int, m2+1)
	}

	for i := 0; i < m1+1; i++ {
		for j := 0; j < m2+1; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) bool
	dfs = func(idx1, idx2 int) bool {
		if idx1 >= m1 && idx2 >= m2 {
			return true
		}

		if dp[idx1][idx2] != -1 {
			return dp[idx1][idx2] == 1
		}

		dp[idx1][idx2] = 0

		k := idx1 + idx2
		if idx1 < m1 && s1[idx1] == s3[k] && dfs(idx1+1, idx2) {
			dp[idx1][idx2] = 1
		}

		if dp[idx1][idx2] == 0 && idx2 < m2 && s2[idx2] == s3[k] && dfs(idx1, idx2+1) {
			dp[idx1][idx2] = 1
		}

		return dp[idx1][idx2] == 1
	}

	return dfs(0, 0)
}
