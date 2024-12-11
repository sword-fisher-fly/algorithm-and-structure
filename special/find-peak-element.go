package special

import (
	"math"
	"math/rand"
)

//https://leetcode.cn/problems/find-peak-element/submissions/584579644/?envType=study-plan-v2&envId=top-interview-150

func FindPeakElement(nums []int) int {
	n := len(nums)

	idx := rand.Intn(n)

	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt
		}

		return nums[i]
	}

	for !(get(idx-1) < get(idx) && get(idx) > get(idx+1)) {
		if get(idx) < get(idx+1) {
			idx++
		} else {
			idx--
		}
	}

	return idx
}
