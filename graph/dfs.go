package graph

func dfs(grid [][]int, visited [][]bool, x, y int) {
	for _, d := range directions {
		nextX, nextY := x+d[0], y+d[1]
		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) || grid[nextX][nextY] == 0 || visited[nextX][nextY] {
			continue
		}

		// if !visited[nextX][nextY] && grid[nextX][nextY] == 1 {
		visited[nextX][nextY] = true
		dfs(grid, visited, nextX, nextY)
		// }
	}
}

func dfsII(grid [][]int, visited [][]bool, x, y int) {
	if visited[x][y] || grid[x][y] == 0 {
		return
	}
	
	visited[x][y] = true

	for _, d := range directions {
		nextX, nextY := x+d[0], y+d[1]
		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
			continue
		}

		dfsII(grid, visited, nextX, nextY)
	}
}
