package binarysearch

//https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/?envType=study-plan-v2&envId=top-interview-150
// 4,5,6,7,0,1,2

// 0,1,2,3,4,5,6
func FindMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[left]
}
