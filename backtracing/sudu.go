package backtracking

import (
	"strings"
)

func isValidSudu(board [][]byte, row, col int, val byte) bool {
	// Scan row
	// for i := 0; i < 9; i++ {
	// 	if board[row][i] == val {
	// 		return false
	// 	}
	// }

	// for i := 0; i < 9; i++ {
	// 	if board[i][col] == val {
	// 		return false
	// 	}
	// }

	// startX, startY := row/3, col/3

	// for i := startX * 3; i < startX*3+3; i++ {
	// 	for j := startY * 3; j < startY*3+3; j++ {
	// 		if board[i][j] == val {
	// 			return false
	// 		}
	// 	}
	// }

	bi, bj := row/3*3, col/3*3

	// 按照数独的规则，检查 b 能否放在 board[r][c]

	for n := 0; n < 9; n++ {
		if board[row][n] == val ||
			board[n][col] == val ||
			board[bi+n/3][bj+n%3] == val {
			return false
		}
	}
	return true

	// return true
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

func SolveSudoku(board [][]byte) {
	solve(board, 0)
}

/* k 是把 board 转换成一维数组后，元素的索引值 */
func solve(board [][]byte, k int) bool {
	if k == 81 {
		return true
	}

	r, c := k/9, k%9
	if board[r][c] != '.' {
		return solve(board, k+1)
	}

	/* bi, bj 是 rc 所在块的左上角元素的索引值 */
	bi, bj := r/3*3, c/3*3

	// 按照数独的规则，检查 b 能否放在 board[r][c]
	isValid := func(b byte) bool {
		for n := 0; n < 9; n++ {
			if board[r][n] == b ||
				board[n][c] == b ||
				board[bi+n/3][bj+n%3] == b {
				return false
			}
		}
		return true
	}

	for b := byte('1'); b <= '9'; b++ {
		if isValid(b) {
			board[r][c] = b
			if solve(board, k+1) {
				return true
			}
		}
	}

	board[r][c] = '.'

	return false
}

func IsValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		if !isValidSudokuRow(board, row) {
			return false
		}
	}

	for col := 0; col < 9; col++ {
		if !isValidSudokuCol(board, col) {
			return false
		}
	}

	for pod := 0; pod < 9; pod++ {
		if !isValidSudokuPod(board, pod) {
			return false
		}
	}
	return true
}

func isValidSudokuRow(board [][]byte, row int) bool {
	var nums [10]bool
	for col := 0; col < 9; col++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidSudokuCol(board [][]byte, col int) bool {
	var nums [10]bool
	for row := 0; row < 9; row++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidSudokuPod(board [][]byte, pod int) bool {
	var nums [10]bool

	row := (pod / 3) * 3
	col := (pod % 3) * 3

	for drow := 0; drow < 3; drow++ {
		for dcol := 0; dcol < 3; dcol++ {
			n := convertToNumber(board[row+drow][col+dcol])
			if n < 0 {
				continue
			}
			if nums[n] {
				return false
			}
			nums[n] = true
		}
	}
	return true
}

func convertToNumber(b byte) int {
	if b == '.' {
		return -1
	}
	return int(b - '0')
}
