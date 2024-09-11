package binarysearch

// [left, right]
func BinarySearchInSortedArray(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// [left, right)

func BinarySearchInSortedArrayII(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)>>1
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return -1
}

func BinarySearchInSortedArrayRecursive(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 && arr[0] != target {
		return -1
	}

	left, right := 0, len(arr)
	mid := left + (right-left)/2
	if target == arr[mid] {
		return mid
	}
	if target < arr[mid] {
		return BinarySearchInSortedArrayRecursive(arr[:mid], target)
	}

	return BinarySearchInSortedArrayRecursive(arr[mid+1:], target)
}
