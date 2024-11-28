package dfs

// import "math"

// func MinPathSum(grid [][]int) int {
// 	minPathSum := math.MaxInt

// 	// right -> down -> left -> up
// 	directions := [4][2]int{
// 		{1, 0},
// 		{0, -1},
// 		{-1, 0},
// 		{0, 1},
// 	}
// 	isInBoard := func(grid [][]int, x, y int) bool {
// 		return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
// 	}

// 	visited := make([][]bool, len(grid))
// 	for i := range visited {
// 		visited[i] = make([]bool, len(grid[0]))
// 	}

// 	var dfs func(int, int, int)
// 	dfs = func(x, y, pathSum int) {
// 		if x == len(grid)-1 && y == len(grid[0])-1 {
// 			if pathSum < minPathSum {
// 				minPathSum = pathSum
// 			}
// 			return
// 		}

// 		for _, d := range directions {
// 			nextX, nextY := d[0]+x, d[1]+y
// 			if isInBoard(grid, nextX, nextY) && !visited[nextX][nextY] {
// 				visited[nextX][nextY] = true
// 				dfs(nextX, nextY, pathSum+grid[nextX][nextY])
// 				visited[nextX][nextY] = false
// 			}
// 		}
// 	}

// 	dfs(0, 0, grid[0][0])

// 	return minPathSum
// }

func MinPathSumByRecursive(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}
	var helper func([][]int, int, int) int
	helper = func(grid [][]int, i, j int) int {
		if i == n-1 && j == m-1 {
			return grid[i][j]
		}

		if i == n-1 {
			return grid[i][j] + helper(grid, i, j+1)
		}

		if j == m-1 {
			return grid[i][j] + helper(grid, i+1, j)
		}

		return grid[i][j] + min(helper(grid, i+1, j), helper(grid, i, j+1))
	}

	return helper(grid, 0, 0)
}
