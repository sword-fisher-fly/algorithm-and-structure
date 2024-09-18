package backtracking

import (
	"fmt"
	"strings"
)

// 1) 扫描列
// 2) 扫描45度 i=row-1,j = col-1
// 3) 扫描135度 i=row-1,j = col+1
func isValidBoard(row, col int, chessboard [][]byte, n int) bool {
	for i := 0; i < n; i++ {
		if chessboard[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func prettyFormatBoard(chessboard [][]byte) string {
	res := strings.Builder{}

	for i := range chessboard {
		for j := range chessboard[i] {
			res.WriteByte(chessboard[i][j])
		}
		res.WriteByte('\n')
	}

	return res.String()
}
func SolveNQueens(n int) [][]string {
	res := make([][]string, 0)
	chessboard := make([][]byte, n)
	for i := range chessboard {
		chessboard[i] = make([]byte, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = '.'
		}
	}

	var backtracking func(n, row int, chessboard [][]byte)
	backtracking = func(n, row int, chessboard [][]byte) {
		if row == n {
			fmt.Printf("chessboard:\n%s\n", prettyFormatBoard(chessboard))
			temp := make([]string, n)
			for i := 0; i < n; i++ {
				temp[i] = string(chessboard[i])
			}
			res = append(res, temp)
			return
		}

		for col := 0; col < n; col++ {
			if !isValidBoard(row, col, chessboard, n) {
				continue
			}
			chessboard[row][col] = 'Q'
			backtracking(n, row+1, chessboard)
			chessboard[row][col] = '.'
		}
	}

	backtracking(n, 0, chessboard)

	return res
}
