package dfs

func GenerateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	var dfs func(int, int, string)
	dfs = func(left, right int, path string) {
		if left == 0 && right == 0 {
			res = append(res, path)
			return
		}

		if left > 0 {
			dfs(left-1, right, path+"(")
		}

		if right > 0 && left < right {
			dfs(left, right-1, path+")")
		}
	}

	dfs(n, n, "")

	return res
}
