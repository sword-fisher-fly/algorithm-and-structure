package array

// 已知一个数组内元素值在0..n-1之间, 可能有多个重复值
// 求: 找出其中一个重复值
// Input: [2,1,3,2,5,4,3] n=7, element: [0, 6]
func FindRepeatedElement(nums []int) int {
	for i, num := range nums {
		for i != num { // If i == num at the first iteration, it can happen the value of the position i is equal to i.
			j := num
			if nums[j] == j {
				return j
			}
			// replace the position of j with the value j
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return -1
}
