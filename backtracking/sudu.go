package backtracking

import (
	"strings"
)

func isValidSudu(board [][]byte, row, col int, val byte) bool {
	// Scan row
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if board[i][col] == val {
			return false
		}
	}

	startX, startY := row/3, col/3

	for i := startX * 3; i < startX*3+3; i++ {
		for j := startY * 3; j < startY*3+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}

	return true
}

func viewBoard(board [][]byte) string {
	ret := strings.Builder{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ret.WriteByte(board[i][j])
			ret.WriteByte(byte(' '))
		}
		ret.WriteString("\n")
	}

	return ret.String()
}
func SolveSudu(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}

			for k := 1; k <= 9; k++ {
				if isValidSudu(board, i, j, byte(k+'0')) {
					board[i][j] = byte(k + '0')
					if SolveSudu(board) {
						return true
					}
					board[i][j] = '.'
				}
			}

			return false
		}
	}

	return true
}
