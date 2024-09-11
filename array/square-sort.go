package array

func SquareSort(nums []int) []int {
	startIdx, endIdx := 0, len(nums)-1

	result := make([]int, len(nums))
	idx := len(nums) - 1
	for startIdx <= endIdx {
		if nums[startIdx]*nums[startIdx] < nums[endIdx]*nums[endIdx] {
			result[idx] = nums[endIdx] * nums[endIdx]
			endIdx--
		} else {
			result[idx] = nums[startIdx] * nums[startIdx]
			startIdx++
		}

		idx--
	}

	return result
}
