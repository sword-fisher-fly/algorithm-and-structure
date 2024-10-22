package binarysearch

// TODO:
// https://leetcode.cn/problems/peak-index-in-a-mountain-array/

// 注意点
// 两种思路
// [left, right): 需要mid一定有变化
// [left, right]: 需要left和right都一定能变化

//  1. 注意边界条件 left:0 right:1 需要让left也能变化，否则一直为0
func PeakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left+1)/2
		if arr[mid] > arr[mid-1] {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

func PeakIndexInMountainArrayII(arr []int) int {
	left, right := 0, len(arr)-2
	for left <= right {
		mid := left + (right-left)/2 // mid > 0 && mid < len(arr)-1
		if arr[mid] > arr[mid+1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
