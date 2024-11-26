package sort

// https://leetcode.cn/problems/kth-largest-element-in-an-array/
func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	// [3, 2, 1, 5, 6, 4]
	// 1)partion = 3, i = 0, j=2, [1,2 3,5,6,4]
	// 2)
	for i < j {
		// 左边 找到第一个 nums[i] >= partion
		i++
		for nums[i] < partition {
			i++
		}

		// 右边找到第一个 nums[j] <= partion
		j--

		for nums[j] > partition {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if k <= j {
		// [1,2 3,5,6,4] l=0, j =2, k=2
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

func FindKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}
