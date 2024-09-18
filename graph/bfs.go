package graph

func bfs(grid [][]int, visited [][]bool, x, y int) {
	var queue [][2]int
	queue = append(queue, [2]int{x, y})
	visited[x][y] = true

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		curX, curY := cur[0], cur[1]

		for _, d := range directions {
			nextX, nextY := curX+d[0], curY+d[1]
			if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
				continue
			}

			if !visited[nextX][nextY] && grid[nextX][nextY] == 1 {
				queue = append(queue, [2]int{nextX, nextY})
				visited[nextX][nextY] = true
			}
		}
	}
}
