package dfs

func MaxPathAscendingInMatrix(matrix [][]int) int {
	// write code here
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	isInBoard := func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n
	}

	directions := [4][2]int{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}
	maxPath := 1
	var dfs func(int, int, int)
	dfs = func(x, y int, depth int) {
		if depth > maxPath {
			maxPath = depth
		}

		for _, d := range directions {
			nextX, nextY := x+d[0], y+d[1]
			if isInBoard(nextX, nextY) && matrix[nextX][nextY] > matrix[x][y] {
				dfs(nextX, nextY, depth+1)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, 1)
		}
	}

	return maxPath
}
