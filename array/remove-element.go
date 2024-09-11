package array

func RemoveElementByForce(nums []int, val int) int {
	arrLen := len(nums)

	for i := 0; i < arrLen; i++ {
		if nums[i] == val {
			for j := i; j < arrLen-1; j++ {
				nums[j] = nums[j+1]
			}

			arrLen--
			i--
		}
	}

	return arrLen
}

func RemoveElementByTwoPointers(nums []int, val int) int {
	slowIdx := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if val != nums[fastIndex] {
			nums[slowIdx] = nums[fastIndex]
			slowIdx++
		}
	}

	return slowIdx
}
