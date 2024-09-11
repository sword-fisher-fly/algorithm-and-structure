package array

func GenerateMatrix(n int) [][]int {
	startx, starty := 0, 0
	var loop int = n / 2
	var center int = n / 2
	count := 1
	offset := 1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for loop > 0 {
		i, j := startx, starty

		//行数不变 列数在变
		for j = starty; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}
		//列数不变是j 行数变
		for i = startx; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		//行数不变i 列数变 j--
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		//列不变j 行变i--
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}

	if n%2 == 1 {
		res[center][center] = n * n
	}
	return res
}
