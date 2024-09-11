package monotonic

// 循环数组: 最后一个元素是数组的第一个元素, 输出每个元素的下一个更大元素
// [1,2,1]  -> [2, -1, 2]
func NextGreaterNumberInCircleArray(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	result := make([]int, len(nums))

	for i := range result {
		result[i] = -1
	}

	stack := []int{0}

	length := len(nums)
	for i := 1; i < length*2; i++ {
		if nums[i%length] <= nums[stack[len(stack)-1]] {
			stack = append(stack, i%length)
		} else {
			for len(stack) > 0 && nums[i%length] > nums[stack[len(stack)-1]] {
				result[stack[len(stack)-1]] = nums[i%length]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i%length)
		}
	}

	return result
}

func NextGreaterNumberInCircleArrayII(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	result := make([]int, len(nums))

	for i := range result {
		result[i] = -1
	}

	stack := []int{0}

	length := len(nums)
	for i := 1; i < length*2; i++ {
		for len(stack) > 0 && nums[i%length] > nums[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = nums[i%length]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%length)

	}

	return result
}
