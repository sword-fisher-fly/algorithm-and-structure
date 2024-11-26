package matrix

// 前提: matrix是一个正方形矩阵
func swapMatrixBySymmetric(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func swapMatrixTopBottom(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
}

func swapMatrixLeftRight(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

// matrix:
//
//	1,2,3
//	4,5,6
//	7,8,9
//
// rotated:
//
//	7,4,1
//	8,5,2
//	9,6,3

// 顺时针旋转90度
//  1. 上下对称交换
//     7,8,9
//     4,5,6
//     1,2,3
//  2. 对角线对称交换
//     7,4,1
//     8,5,2
//     9,6,3
func RotateByClockwise(matrix [][]int) {
	swapMatrixTopBottom(matrix)
	swapMatrixBySymmetric(matrix)
}

// 逆时针旋转90度
//  1. 左右对称交换
//     3,2,1
//     6,5,4
//     9,8,7
//  2. 对角线对称交换
//     3,6,9
//     2,5,8
//     1,4,7
//     再逆时针旋转90度
//     9,8,7
//     6,5,4
//     3,2,1
func RotateByAntiClockwise(matrix [][]int) {
	swapMatrixLeftRight(matrix)
	swapMatrixBySymmetric(matrix)
}

// 顺时针旋转180度
//  1. 左右对称交换
//     3,2,1
//     6,5,4
//     9,8,7
//  2. 上下对称交换
//     9,8,7
//     6,5,4
//     3,2,1
func RotateBy180Clockwise(matrix [][]int) {
	swapMatrixLeftRight(matrix)
	swapMatrixTopBottom(matrix)
}

func RotateBy180ClockwiseII(matrix [][]int) {
	RotateByClockwise(matrix)
	RotateByClockwise(matrix)
}

// 逆时针旋转180度
//  1. 上下对称交换
//     7,8,9
//     4,5,6
//     1,2,3
//  2. 左右对称交换
//     9,8,7
//     6,5,4
//     3,2,1
func RotateBy180AntiClockwise(matrix [][]int) {
	swapMatrixLeftRight(matrix)
	swapMatrixTopBottom(matrix)
}
