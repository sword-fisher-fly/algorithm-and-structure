package sort

import "fmt"

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
		fmt.Printf("nums[i]=%d, partition=%d\n", nums[i], partition)
		for nums[i] < partition {
			i++
		}

		// 右边找到第一个 nums[j] <= partion
		j--
		fmt.Printf("nums[j]=%d, partition=%d\n", nums[j], partition)

		for nums[j] > partition {
			j--
		}

		fmt.Printf("swap: i=%d, j=%d\n", i, j)
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if k <= j {
		// [1,2 3,5,6,4] l=0, j =2, k=2
		fmt.Printf("quickselect(nums, l, j, k): l=%d, j =%d, k=%d\n", l, j, k)
		return quickselect(nums, l, j, k)
	} else {
		fmt.Printf("quickselect(nums, j+1, r, k): j+1=%d,r=%d,k=%d\n", j+1, r, k)
		return quickselect(nums, j+1, r, k)
	}
}

func FindKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}
