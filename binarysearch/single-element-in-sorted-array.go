package binarysearch

// TODO:
// https://leetcode.cn/problems/single-element-in-a-sorted-array/description/

// [1,1,2,3,3,4,4,8,8]
// 1) leftIdx: 0, rightIdx: 8, midIdx: 4
// 2) leftIdx: 0, rightIdx: 4, midIdx: 2
// 3) leftIdx: 0, rightIdx: 2, midIdx: 1
// 4) leftIdx: 2, rightIdx: 2, midIdx: 2
func SingleElementInSortedArray(nums []int) int {
	leftIdx := 0
	rightIdx := len(nums) - 1
	for leftIdx < rightIdx {
		mid := leftIdx + (rightIdx-leftIdx)/2
		if nums[mid] != nums[mid^1] {
			rightIdx = mid
		} else {
			leftIdx = mid + 1
		}
	}

	return nums[leftIdx]
}
