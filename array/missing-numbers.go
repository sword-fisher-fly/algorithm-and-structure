package array

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

// in-place method
// Input: [4,3,2,7,8,2,3,1]
// Output: [5,6]
// 3, 4, 3, 1, 5
func FindMissingNumbers(nums []int) []int {
	for _, num := range nums {
		pos := abs(num) - 1
		if nums[pos] > 0 {
			nums[pos] = -nums[pos]
		}
	}

	result := []int{}
	for i := range nums {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}
