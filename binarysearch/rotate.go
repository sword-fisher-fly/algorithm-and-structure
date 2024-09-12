package binarysearch

// rotated array: [6,7,1,2,3,4,5], target: 3
func BinarySearchInRotatedArray(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[left] < arr[mid] { // 左边有序
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右边有序
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
