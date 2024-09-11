package array

// 数组arr, 计算任意区间的元素和arr[i:j] i>=0, j<=n-1
// i < j
func IntervalSum(arr []int, i, j int) int {
	if len(arr) == 0 {
		return -1
	}

	if i < 0 || j > len(arr) {
		return -1
	}

	if i > j {
		return -1
	}

	// 记pSum[i-1]= arr[0]+...+arr[i-1]
	pSum := make([]int, len(arr))
	pSum[0] = arr[0]
	for k := 1; k < len(arr); k++ {
		pSum[k] = pSum[k-1] + arr[k]
	}

	if i == 0 {
		return pSum[j-1]
	}

	if i == j {
		return arr[i]
	}

	return pSum[j-1] - pSum[i-1]
}
