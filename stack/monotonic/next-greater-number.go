package monotonic

// nums1 是 nums2的子集
// 返回 nums1 中每个元素在 nums2 中下一个比其大的数
func NextGreaterNumber(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))

	for i := range result {
		result[i] = -1
	}

	// key:val -> nums[i]:i
	nums1Map := map[int]int{}

	for i := range nums1 {
		nums1Map[nums1[i]] = i
	}

	stack := []int{0}
	for i := 1; i < len(nums2); i++ {
		if nums2[i] <= nums2[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] {
				if index, ok := nums1Map[nums2[stack[len(stack)-1]]]; ok {
					result[index] = nums2[i]
				}
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, i)
		}
	}

	return result
}

func NextGreaterElementII(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}
	m := make(map[int]int, len(nums1))
	for k, v := range nums1 {
		m[v] = k
	}

	stack := []int{0}
	for i := 1; i < len(nums2); i++ {
		top := stack[len(stack)-1]
		if nums2[i] < nums2[top] {
			stack = append(stack, i)
		} else if nums2[i] == nums2[top] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && nums2[i] > nums2[top] {
				if v, ok := m[nums2[top]]; ok {
					res[v] = nums2[i]
				}
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}
