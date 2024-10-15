package graph

// 岛屿问题
// 1) 岛屿数量
// 2) 岛屿周长
// 3) 最大岛屿面积
// 4) 最大岛屿面积增强版
// 5) 孤岛总面积/沉没孤岛(同等问题)

// 1: 陆地 0: 水域
// 输出: 计算岛屿的数量
// 说明: 岛屿由水平方向和垂直方向上相邻陆地连接而成,并且四周都是水域
// 假设矩阵外都被水包围
/*
1 1 0 0 0
1 1 0 0 0
0 0 1 0 0
0 0 0 1 1
*/

// 1) 岛屿数量
func FindIslandsByBFS(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	result := 0
	bfs := func(grid [][]int, visited [][]bool, x, y int) {
		queue := [][2]int{{x, y}}
		visited[x][y] = true
		for len(queue) != 0 {
			curX, curY := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range directions {
				nextX, nextY := curX+d[0], curY+d[1]
				if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || grid[nextX][nextY] == 0 || visited[nextX][nextY] {
					continue
				}

				visited[nextX][nextY] = true
				queue = append(queue, [2]int{nextX, nextY})
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				bfs(grid, visited, i, j)
				result++
			}
		}
	}

	return result
}

func FindIslandsByDFS(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var result int
	var dfs func([][]int, [][]bool, int, int)

	dfs = func(grid [][]int, visited [][]bool, x, y int) {
		for _, d := range directions {
			nextX, nextY := x+d[0], y+d[1]
			if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || visited[nextX][nextY] || grid[nextX][nextY] == 0 {
				continue
			}

			visited[nextX][nextY] = true
			dfs(grid, visited, nextX, nextY)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				visited[i][j] = true
				dfs(grid, visited, i, j)
				result++
			}
		}
	}

	return result
}

// 3) 最大岛屿面积
func MaxIslandArea(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	maxArea := 0

	count := 0
	var dfs func([][]int, [][]bool, int, int)
	dfs = func(grid [][]int, visited [][]bool, x, y int) {
		for _, d := range directions {
			nextX, nextY := x+d[0], y+d[1]
			if isInBoard(grid, nextX, nextY) && !visited[nextX][nextY] && grid[nextX][nextY] == 1 {
				visited[nextX][nextY] = true
				count++
				dfs(grid, visited, nextX, nextY)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				visited[i][j] = true
				count = 1
				dfs(grid, visited, i, j)
				if count > maxArea {
					maxArea = count
				}
			}
		}
	}

	return maxArea
}

// 4) 最大岛屿面积增强版, 通过最多人工改变一个0为1后的最大岛屿面积
func FindLargestAreaIslandByBFS(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// 借助mark标记相连的陆地
	mark := 1
	gridCountMap := make(map[int]int)

	isAllLand := true
	var bfs func([][]int, int, int)
	bfs = func(grid [][]int, x, y int) {
		grid[x][y] = mark
		queue := [][2]int{{x, y}}

		gridCountMap[mark] = 1

		for len(queue) > 0 {
			curX, curY := queue[0][0], queue[0][1]
			queue = queue[1:]

			for _, d := range directions {
				nextX, nextY := curX+d[0], curY+d[1]
				if isInBoard(grid, nextX, nextY) && !visited[nextX][nextY] && grid[nextX][nextY] == 1 {
					grid[nextX][nextY] = mark
					gridCountMap[mark]++
					queue = append(queue, [2]int{nextX, nextY})
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				isAllLand = false
			}

			if !visited[i][j] && grid[i][j] == 1 {
				mark++
				bfs(grid, i, j)
			}
		}
	}

	if isAllLand {
		return m * n
	}

	maxArea := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count := 1
				for _, d := range directions {
					nextX, nextY := i+d[0], j+d[1]
					if isInBoard(grid, nextX, nextY) && grid[nextX][nextY] != 0 {
						count += gridCountMap[grid[nextX][nextY]]
					}
				}

				if count > maxArea {
					maxArea = count
				}
			}
		}
	}

	return maxArea
}

func FindLargestAreaIslandByDFS(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	mark := 2

	isAllLands := true

	gridCountMap := make(map[int]int)

	var dfs func(int, int)
	dfs = func(x, y int) {
		grid[x][y] = mark
		gridCountMap[mark]++

		for _, d := range directions {
			nextX, nextY := x+d[0], y+d[1]
			if isInBoard(grid, nextX, nextY) && !visited[nextX][nextY] && grid[nextX][nextY] == 1 {
				dfs(nextX, nextY)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				isAllLands = false
			}

			if !visited[i][j] && grid[i][j] == 1 {
				dfs(i, j)
				mark++
			}
		}
	}

	if isAllLands {
		return m * n
	}

	maxArea := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count := 1
				for _, d := range directions {
					nextX, nextY := i+d[0], j+d[1]
					if isInBoard(grid, nextX, nextY) && grid[nextX][nextY] != 0 {
						count += gridCountMap[grid[nextX][nextY]]
					}
				}

				if count > maxArea {
					maxArea = count
				}
			}
		}
	}

	return maxArea
}

// 5) 孤岛总面积
// 孤岛是那些位于矩阵内部、所有单元格都不接触边缘的岛屿
func MaxIsolatedIslandAreaByDFS(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	var dfs func([][]int, int, int)
	dfs = func(grid [][]int, x, y int) {
		grid[x][y] = 0
		for _, d := range directions {
			nextX, nextY := x+d[0], y+d[1]
			if isInBoard(grid, nextX, nextY) && grid[nextX][nextY] == 1 {
				dfs(grid, nextX, nextY)
			}
		}
	}

	// 上下边
	for i := 0; i < n; i++ {
		if grid[0][i] == 1 {
			dfs(grid, 0, i)
		}

		if grid[m-1][i] == 1 {
			dfs(grid, m-1, i)
		}
	}
	// 左右边
	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			dfs(grid, i, 0)
		}

		if grid[i][n-1] == 1 {
			dfs(grid, i, n-1)
		}
	}

	area := 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 {
				area++
			}
		}
	}

	return area
}

func MaxIsolatedIslandAreaByBFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var bfs func([][]int, int, int)
	bfs = func(grid [][]int, x, y int) {
		grid[x][y] = 0
		queue := [][2]int{{x, y}}
		for len(queue) > 0 {
			curX, curY := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range directions {
				nextX, nextY := curX+d[0], curY+d[1]
				if isInBoard(grid, nextX, nextY) && grid[nextX][nextY] == 1 {
					grid[nextX][nextY] = 0
					queue = append(queue, [2]int{nextX, nextY})
				}
			}
		}
	}

	// 上下边界扫描
	for i := 0; i < n; i++ {
		if grid[0][i] == 1 {
			bfs(grid, 0, i)
		}

		if grid[m-1][i] == 1 {
			bfs(grid, m-1, i)
		}
	}

	// 左右边界扫描
	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			bfs(grid, i, 0)
		}

		if grid[i][n-1] == 1 {
			bfs(grid, i, n-1)
		}
	}

	area := 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 {
				area++
			}
		}
	}

	return area
}
