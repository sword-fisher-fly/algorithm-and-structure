package backtracking

func Permutation(s string) []string {
	result := make([]string, 0)
	used := make([]bool, len(s))

	path := make([]rune, 0)
	var backTracing func(string, []rune, []bool)

	backTracing = func(s string, path []rune, used []bool) {
		if len(path) == len(s) {
			result = append(result, string(path))
			return
		}

		for i := 0; i < len(s); i++ {
			if used[i] {
				continue
			}

			path = append(path, rune(s[i]))
			used[i] = true
			backTracing(s, path, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backTracing(s, path, used)

	return result
}

// [1,2,2]
// [1,2,2], [2,1,2] [2,2,1]
func PermutationWithDuplicates(nums []int) [][]int {
	result := [][]int{}

	used := make([]bool, len(nums))

	var backTracing func([]int, []int)
	backTracing = func(nums []int, path []int) {
		if len(path) == len(nums) {
			// temp := make([]int, len(path))
			// copy(temp, path)
			result = append(result, path)
			return
		}

		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backTracing(nums, path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backTracing(nums, []int{})

	return result
}
