package offer

// 输入: [4, 2, 5, 7]
// 输出: [4,7,2,5] or [2,5,4,7] or [2,7,4,5]
func SortArrayByParity(nums []int) []int {
	oddIndex := 1
	for i := 0; i < len(nums); i = i + 2 {
		if nums[i]%2 == 1 {
			for oddIndex < len(nums) && nums[oddIndex]%2 == 1 {
				oddIndex += 2
			}

			// 保护边界
			if oddIndex >= len(nums) {
				break
			}

			nums[i], nums[oddIndex] = nums[oddIndex], nums[i]
		}
	}

	return nums
}

func SortArrayByParityII(nums []int) []int {
	ans := make([]int, len(nums))

	oddIndex := 1
	evenIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			ans[oddIndex] = nums[i]
			oddIndex += 2
		} else {
			ans[evenIndex] = nums[i]
			evenIndex += 2
		}
	}

	return ans
}

// nums数量必须为偶数
// func SortArrayByParityII(nums []int) []int {
// 	left, right := 0, len(nums)-1

// 	for left < right {
// 		for left < right && nums[left]%2 == 0 {
// 			left += 2
// 		}

// 		for left < right && nums[right]%2 == 1 {
// 			right -= 2
// 		}

// 		if left >= right {
// 			break
// 		}

// 		nums[left], nums[right] = nums[right], nums[left]
// 	}

// 	return nums
// }
