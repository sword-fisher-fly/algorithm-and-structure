package math

//  下一个排列
// 1,2,4,3 -> 1,3,2,4
// 1) 从后往前找到第一个升序的位置
// 2) 交换
// 3) 翻转

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func NextPermutation(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	for i := len(nums) - 1; i > 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				reverse(nums[i+1:])
				return nums
			}
		}
	}

	reverse(nums)

	return nums
}
