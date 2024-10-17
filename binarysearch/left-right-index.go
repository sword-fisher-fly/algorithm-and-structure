package binarysearch

// TODO:

// 1) nums=[1,2,3,3,3,4,5], target=3
// left: 2, right: 4
// 2) nums=[1,2,3,3,3,4,5], target=6
// left: -1, right: -1
// 3) nums=[1,2,3,3,3,4,5], target=1
// left: 0, right: 0
func SearchLeftRightIndex(nums []int, target int) (int, int) {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			left, right = mid, mid
			for left-1 >= 0 && nums[left-1] == target {
				left--
			}

			for right+1 < len(nums) && nums[right+1] == target {
				right++
			}

			return left, right
		}
	}

	return -1, -1
}

// [left,right)
func SearchLeftRightIndexII(nums []int, target int) (int, int) {
	left, right := 0, len(nums)
	found := false
	mid := 0

	for !found && left < right {
		mid = left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] == target {
			found = true
		}
	}

	if found {
		left, right = mid, mid
		for left-1 >= 0 && nums[left-1] == target {
			left--
		}

		for right+1 < len(nums) && nums[right+1] == target {
			right++
		}

		return left, right
	}

	return -1, -1
}
