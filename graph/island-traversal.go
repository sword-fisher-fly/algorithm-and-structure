package graph

// TODO:
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

				// if !visited[nextX][nextY] {
				visited[nextX][nextY] = true
				queue = append(queue, [2]int{nextX, nextY})
				// visited[nextX][nextY] = true
				// }
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
