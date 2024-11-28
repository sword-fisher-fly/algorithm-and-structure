package sort

import (
	"math"
	"sort"
)

func FindUnsortedSubarray(nums []int) int {
	n := len(nums)
	maxn := math.MinInt
	minn := math.MaxInt
	left, right := -1, -1
	for i := 0; i < n; i++ {
		if maxn > nums[i] {
			right = i
		} else {
			maxn = nums[i]
		}

		if minn < nums[n-1-i] {
			left = n - 1 - i
		} else {
			minn = nums[n-1-i]
		}
	}

	if right == -1 {
		return 0
	}

	return right - left + 1
}

func FindUnSortedSubarrayII(nums []int) int {
	isSorted := func(arr []int) bool {
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				return false
			}
		}

		return true
	}

	if isSorted(nums) {
		return 0
	}

	sortedNum := make([]int, len(nums))
	copy(sortedNum, nums)

	sort.Ints(sortedNum)

	left, right := 0, len(nums)-1
	for nums[left] == sortedNum[left] {
		left++
	}

	for nums[right] == sortedNum[right] {
		right--
	}

	return right - left + 1
}
