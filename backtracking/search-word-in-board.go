package backtracking

// board:
//
//	 [
//	     ['A', 'B', 'C'],
//		 ['C', 'D', 'E'],
//	     ['F', 'G', 'H']
//	 ]
func SearchWordInBoard(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	choosePath := make([][]bool, m)
	for i := range choosePath {
		choosePath[i] = make([]bool, n)
	}

	// 左 -> 下 -> 右 -> 上
	direction := [4][2]int{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}

	var backtracking func(board [][]byte, pathIndex int, word string, i, j int) bool

	backtracking = func(board [][]byte, pathIndex int, word string, i, j int) bool {
		if pathIndex == len(word) {
			return true
		}

		for _, d := range direction {
			nextX := i + d[0]
			nextY := j + d[1]

			if nextX >= 0 && nextX < m && nextY >= 0 && nextY < n && !choosePath[nextX][nextY] && board[nextX][nextY] == word[pathIndex] {
				choosePath[nextX][nextY] = true
				if backtracking(board, pathIndex+1, word, nextX, nextY) {
					return true
				}
				choosePath[nextX][nextY] = false
			}
		}

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word[0] != board[i][j] {
				continue
			}

			choosePath[i][j] = true
			if backtracking(board, 1, word, i, j) {
				return true
			}
			choosePath[i][j] = false
		}
	}

	return false
}
