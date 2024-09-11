package backtracking

func subsetsHelper(arr []int, startIndex int, path []int, result *[][]int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*result = append(*result, tmp)
	if startIndex == len(arr) {
		return
	}

	for i := startIndex; i < len(arr); i++ {
		path = append(path, arr[i])
		subsetsHelper(arr, i+1, path, result)
		path = path[:len(path)-1]
	}
}
func Subsets(arr []int) [][]int {
	result := make([][]int, 0)
	path := []int{}

	subsetsHelper(arr, 0, path, &result)

	return result
}

func subsetsWithDupHelper(arr []int, startIndex int, used *[]bool, path []int, result *[][]int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*result = append(*result, tmp)
	// if startIndex == len(arr) {
	// 	return
	// }

	for i := startIndex; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] && (*used)[i-1] == false {
			continue
		}

		(*used)[i] = true
		path = append(path, arr[i])
		subsetsWithDupHelper(arr, i+1, used, path, result)
		(*used)[i] = false
		path = path[:len(path)-1]
	}
}
func SubsetsWithDup(arr []int) [][]int {
	result := make([][]int, 0)
	path := []int{}
	used := make([]bool, len(arr))

	subsetsWithDupHelper(arr, 0, &used, path, &result)

	return result
}

func subsetsWithDupIIHelper(arr []int, startIndex int, path []int, result *[][]int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*result = append(*result, tmp)

	for i := startIndex; i < len(arr); i++ {
		if i > startIndex && arr[i] == arr[i-1] {
			continue
		}

		path = append(path, arr[i])
		subsetsWithDupIIHelper(arr, i+1, path, result)
		path = path[:len(path)-1]
	}
}
func SubsetsWithDupII(arr []int) [][]int {
	result := make([][]int, 0)
	path := []int{}

	subsetsWithDupIIHelper(arr, 0, path, &result)

	return result
}
