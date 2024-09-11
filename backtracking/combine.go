package backtracking

func combineHelper(n, k int, startIndex int, path []int, result *[][]int) {
	if k == len(path) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for i := startIndex; i <= n; i++ {
		path = append(path, i)
		combineHelper(n, k, i+1, path, result)
		path = path[:len(path)-1]
	}
}
func Combine(n, k int) [][]int {
	result := make([][]int, 0)
	path := []int{}

	combineHelper(n, k, 1, path, &result)

	return result
}

func combineIIHelper(n, k int, startIndex int, path []int, result *[][]int) {
	if k == len(path) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for i := startIndex; i <= n-(k-len(path)); i++ {
		path = append(path, i)
		combineHelper(n, k, i+1, path, result)
		path = path[:len(path)-1]
	}
}
func CombineII(n, k int) [][]int {
	result := make([][]int, 0)
	path := []int{}

	combineHelper(n, k, 1, path, &result)

	return result
}
