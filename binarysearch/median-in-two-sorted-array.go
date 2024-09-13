package binarysearch

import "math"

func mergeTwoSortedArray(nums1, nums2 []int) []int {
	result := make([]int, len(nums1)+len(nums2))
	idx1, idx2 := 0, 0
	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] < nums2[idx2] {
			result[idx1+idx2] = nums1[idx1]
			idx1++
		} else {
			result[idx1+idx2] = nums2[idx2]
			idx2++
		}
	}

	for idx1 < len(nums1) {
		result[idx1+idx2] = nums1[idx1]
		idx1++
	}

	for idx2 < len(nums2) {
		result[idx1+idx2] = nums2[idx2]
		idx2++
	}

	return result
}

func FindMedianInTwoSortedArrayByForce(nums1, nums2 []int) int {
	mergedArray := mergeTwoSortedArray(nums1, nums2)
	if len(mergedArray) == 0 {
		return -1
	}

	if len(mergedArray)%2 == 1 {
		return mergedArray[len(mergedArray)/2]
	}

	return (mergedArray[len(mergedArray)/2] + mergedArray[len(mergedArray)/2-1]) / 2
}

// https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
func findKth(nums1 []int, i int, nums2 []int, j, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}

	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	midVa11 := math.MaxInt
	midVal2 := math.MaxInt
	if i+k/2-1 < len(nums1) {
		midVa11 = nums1[i+k/2-1]
	}

	if j+k/2-1 < len(nums2) {
		midVal2 = nums2[j+k/2-1]
	}

	if midVa11 < midVal2 {
		return findKth(nums1, i+k/2, nums2, j, k-k/2)
	} else {
		return findKth(nums1, i, nums2, j+k/2, k-k/2)
	}
}

func FindMediaInTwoSortedArrayII(nums1, nums2 []int) int {
	m, n := len(nums1), len(nums2)

	left, right := (m+n+1)/2, (m+n+2)/2

	return (findKth(nums1, 0, nums2, 0, left) + findKth(nums1, 0, nums2, 0, right)) / 2
}

// Two Pointer

func FindMedianInTwoSortedArrayByTwoPointer(nums1, nums2 []int) int {

	return -1
}
