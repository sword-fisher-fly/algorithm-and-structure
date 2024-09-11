package twopointer

import "sort"

func FourSum(nums []int) [][4]int {
	if len(nums) < 4 {
		return [][4]int{}
	}

	sort.Ints(nums)
	res := make([][4]int, 0)

	for i := 0; i < len(nums)-3; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for k := i + 1; k < len(nums)-2; k++ {
			if k > i+1 && nums[k] == nums[k-1] {
				continue
			}
			left := k + 1
			right := len(nums) - 1
			for left < right {
				if nums[i]+nums[k]+nums[left]+nums[right] > 0 {
					right--
				} else if nums[i]+nums[k]+nums[left]+nums[right] < 0 {
					left++
				} else {
					res = append(res, [4]int{nums[i], nums[k], nums[left], nums[right]})

					for left < right && nums[left] == nums[left+1] {
						left++
					}

					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				}
			}
		}
	}

	return res
}
