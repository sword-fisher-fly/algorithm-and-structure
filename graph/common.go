package graph

func isInBoard(board [][]int, i, j int) bool {
	return i >= 0 && i < len(board) && j >= 0 && j < len(board[0])
}
