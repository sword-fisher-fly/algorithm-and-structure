package offer

import (
	"fmt"
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
		fmt.Printf("i=%d, x=%d, nums[%d]=%d\n", i, x, i, nums[i])

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
