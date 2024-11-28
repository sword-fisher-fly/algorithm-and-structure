package special

//https://leetcode.cn/problems/maximal-rectangle/description/?envType=problem-list-v2&envId=2cktkvj
func MaximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	left := make([][]int, m)
	for i := range left {
		left[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if j == 0 {
					left[i][j] = 1
				} else {
					left[i][j] = left[i][j-1] + 1
				}
			}
		}
	}

	res := 0

	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			width := left[i][j]
			area := width

			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}

			res = max(res, area)
		}
	}

	return res
}
