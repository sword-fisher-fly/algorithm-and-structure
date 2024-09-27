package backtracking

// 给定数组nums, 找出所有长度大于等于2的递增子序列
// Input: [4, 6, 7, 7]
// Output: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7, 7]]
func FindAscendingSubsequences(nums []int) [][]int {
	res := make([][]int, 0)

	var backtracking func(nums []int, start int, path []int)

	backtracking = func(nums []int, startIndex int, path []int) {
		if len(path) >= 2 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)

			// no return
		}

		// 记录本层已访问过的数值
		set := make(map[int]struct{})
		for i := startIndex; i < len(nums); i++ {
			if len(path) > 0 && nums[i] <= path[len(path)-1] {
				continue
			}

			if _, ok := set[nums[i]]; ok {
				continue
			} else {
				set[nums[i]] = struct{}{}
			}

			path = append(path, nums[i])
			backtracking(nums, i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtracking(nums, 0, []int{})

	return res
}
