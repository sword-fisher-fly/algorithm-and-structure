package binarysearch

// 注意: 审题利用规律

// 每一行从左到右递增, 下一行的第一个元素比上一行的最后一个元素大
// 需要一个目标值是否在数组中, 在则输出[行,列]坐标, 不在则输出[-1,-1]
func FindTargetInSortedMatrix(matrix [][]int, target int) [2]int {
	res := [2]int{-1, -1}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}

	// m=5, n=4, mid=6 -> row = mid/n, col = mid%n
	m, n := len(matrix), len(matrix[0])
	low, high := 0, m*n-1
	for low <= high {

		mid := low + (high-low)>>1
		row, col := mid/n, mid%n
		if matrix[row][col] == target {
			res = [2]int{row, col}
			return res
		} else if matrix[row][col] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return res
}

// 每一行从左到右递增, 每一列从上到下递增
// 利用规律: 从第一行最后一列的元素开始比较, 逐步变更行列
func FindTargetInSortedMatrixII(matrix [][]int, target int) [2]int {
	res := [2]int{-1, -1}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}

	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return [2]int{row, col}
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}

	return [2]int{-1, -1}
}
