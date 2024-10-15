package hard

// 1) 同一列无重复的皇后
// 2) 45度无重复的皇后
// 3) 135度无重复的皇后
// 同一行无重复的皇后(逐行遍历则保证同一行不会出现2个皇后)
func isValidQueen(board [][]string, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	return true
}

func TotalNQueens(n int) int {
	count := 0

	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	var backtracking func(row int)
	backtracking = func(row int) {
		if n == row {
			count++
			return
		}
		for i := 0; i < n; i++ {
			if !isValidQueen(board, row, i) {
				continue
			}

			board[row][i] = "Q"
			backtracking(row + 1)
			board[row][i] = "."
		}
	}

	backtracking(0)

	return count
}
