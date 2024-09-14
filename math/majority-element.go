package math

// MajorityElement returns the majority element in the given slice.
// 投票算法: 假设数组内某个元素出现次数超过一半
func MajorityElement(nums []int) int {
	var count int
	var candidate int

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
