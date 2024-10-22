package offer

// TODO:
// https://leetcode.cn/problems/generate-parentheses/submissions/574826268/

func GenerateParenthesis(n int) []string {
	var dfs func(int, int, string)

	ans := []string{}
	dfs = func(left, right int, path string) {
		if left == n && right == n {
			ans = append(ans, path)
			return
		}

		if left < n {
			dfs(left+1, right, path+"(")
		}
		if right < left {
			dfs(left, right+1, path+")")
		}
	}

	dfs(0, 0, "")
	return ans
}

func GenerateParenthesisII(n int) []string {
	maxLen := n * 2
	ans := []string{}
	var dfs func(int, int, string)
	dfs = func(open, close int, path string) {
		if len(path) == maxLen {

			ans = append(ans, path)
			return
		}

		if open < n {
			dfs(open+1, close, path+"(")
		}

		if close < open {
			dfs(open, close+1, path+")")
		}
	}

	dfs(0, 0, "")

	return ans
}

// 第二种解法leetcode题解

func GenerateParenthesisIII(n int) []string {
	cache := map[int][]string{}
	var generate func(int) []string

	generate = func(n int) []string {
		if cache[n] != nil {
			return cache[n]
		}

		if n == 0 {
			return []string{""}
		}

		ans := []string{}
		for i := 0; i < n; i++ {
			for _, left := range generate(i) {
				for _, right := range generate(n - 1 - i) {
					ans = append(ans, "("+left+")"+right)
				}
			}
		}

		cache[n] = ans
		return ans
	}

	res := generate(n)

	return res
}
