package twopointer

import "sort"

// 给定数组nums，找出符合a+b+c=0条件的3个元素，返回所有满足条件的三元组。
// [-1,0,1,2,-1,-4]
func ThreeSum(nums []int) [][3]int {
	sort.Ints(nums)

	res := [][3]int{}

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, [3]int{nums[i], nums[left], nums[right]})
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

	return res
}
