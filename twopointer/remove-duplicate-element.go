package twopointer

// 1) 移除相同元素，只保留1个
// 2) 移除相同元素，保留2个

func RemoveDuplicate(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	slow, fast := 1, 1
	for fast < n {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}

	return slow
}

func RemoveDuplicateII(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}

	return slow
}
