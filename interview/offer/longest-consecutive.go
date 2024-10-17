package offer

import (
	"sort"
)

func LongestConsecutive(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	sort.Ints(nums)
	ans, t := 1, 1
	for i, x := range nums[1:] {
		if x == nums[i] {
			continue
		}
		if x == nums[i]+1 {
			t++
			ans = max(ans, t)
		} else {
			t = 1
		}
	}

	return ans
}

func LongestConsecutiveII(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	sort.Ints(nums)
	ans, t := 1, 1
	for i := 1; i < n; i++ {
		if nums[i-1] == nums[i] {
			continue
		}

		if nums[i] == nums[i-1]+1 {
			t++
			ans = max(ans, t)
		} else {
			t = 1
		}
	}

	return ans
}
