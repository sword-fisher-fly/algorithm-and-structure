package interview

import (
	"fmt"
	"strings"
)

func colorPathInBoard(board [][]byte, visited [][]bool) string {
	res := strings.Builder{}

	for i := range board {
		for j := range board[i] {
			if visited[i][j] {
				res.WriteString(fmt.Sprintf("%s%c%s", RedColor, board[i][j], ResetColor))
			} else {
				res.WriteByte(board[i][j])
			}
		}
		res.WriteByte('\n')
	}

	return res.String()
}

func SearchWordInBoardByBFS(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])

	bfs := func(board [][]byte, word string, wordIndex int, x, y int) bool {
		if wordIndex == len(word) {
			return true
		}

		visited := make([][]bool, m)
		for i := 0; i < m; i++ {
			visited[i] = make([]bool, n)
		}

		visited[x][y] = true

		queue := [][2]int{{x, y}}

		for len(queue) > 0 {
			curX, curY := queue[0][0], queue[0][1]
			queue = queue[1:]

			for _, dir := range directions {
				nextX, nextY := curX+dir[0], curY+dir[1]
				if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || visited[nextX][nextY] || board[nextX][nextY] != word[wordIndex] {
					continue
				}

				// if board[nextX][nextY] == word[wordIndex] && !visited[nextX][nextY] {
				wordIndex++
				if wordIndex == len(word) {
					visited[nextX][nextY] = true
					fmt.Printf("mark the path:\n%s\n", colorPathInBoard(board, visited))
					return true
				}

				queue = append(queue, [2]int{nextX, nextY})
				visited[nextX][nextY] = true
				break // important here, otherwise it will fail.
				// }
			}
		}

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word[0] != board[i][j] {
				continue
			}

			if bfs(board, word, 1, i, j) {
				return true
			}
		}
	}

	return false
}

func SearchWordInBoardByBacktracing(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	var backtracing func(board [][]byte, word string, wordIndex int, x, y int, visited [][]bool) bool

	backtracing = func(board [][]byte, word string, wordIndex int, x, y int, visited [][]bool) bool {
		if wordIndex == len(word) {
			return true
		}

		for _, d := range directions {
			nextX, nextY := x+d[0], y+d[1]
			if nextX < 0 || nextX >= len(board) || nextY < 0 || nextY >= len(board[0]) || visited[nextX][nextY] || board[nextX][nextY] != word[wordIndex] {
				continue
			}

			visited[nextX][nextY] = true
			if backtracing(board, word, wordIndex+1, nextX, nextY, visited) {
				return true
			}
			visited[nextX][nextY] = false
		}

		return false
	}

	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] != board[i][j] {
				continue
			}

			visited[i][j] = true
			if backtracing(board, word, 1, i, j, visited) {
				return true
			}
			visited[i][j] = false
		}
	}

	return false
}

// TODO: 回溯就是深搜一种, 如从图的角度是否还存在另一种形式的写法??
func SearchWordInBoardByDFS(board [][]byte, word string) bool {

	return false
}
