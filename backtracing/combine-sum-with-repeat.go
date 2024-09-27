package backtracking

// 假设candidates是正整数(不能有0元素)
func combineSumWithRepeat(candidates []int, target int, sum int, startIndex int, result *[][]int, path []int) {
	if sum > target {
		return
	}

	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		sum += candidates[i]
		path = append(path, candidates[i])
		combineSumWithRepeat(candidates, target, sum, i, result, path)
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}
func CombineSumWithRepeat(candidates []int, target int) [][]int {
	result := [][]int{}
	path := []int{}

	combineSumWithRepeat(candidates, target, 0, 0, &result, path)

	return result
}

func combineSumWithRepeatOptimized(candidates []int, target, sum int, startIndex int, result *[][]int, path []int) {
	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		sum += candidates[i]
		path = append(path, candidates[i])
		combineSumWithRepeatOptimized(candidates, target, sum, i, result, path)
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}
func CombineSumWithRepeatOptimized(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	path := []int{}

	combineSumWithRepeatOptimized(candidates, target, 0, 0, &result, path)

	return result
}

func combineSumWithRepeatII(candidates []int, target int, startIndex int, result *[][]int, path []int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		if target < candidates[i] {
			continue
		}

		path = append(path, candidates[i])
		combineSumWithRepeatII(candidates, target-candidates[i], i, result, path)
		path = path[:len(path)-1]
	}
}

func CombineSumWithRepeatII(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}

	combineSumWithRepeatII(candidates, target, 0, &res, path)

	return res
}
