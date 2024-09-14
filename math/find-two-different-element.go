package math

// 给定数组只有两个元素出现一次，其他元素出现两次，找出这两个元素

// num^num =0
// num&(num-1) 消除最后一个二进制位为1的位
// (num&(num-1))^num 取最后一个进制位为1的位
func FindTwoDifferentElement(nums []int) [2]int {
	if len(nums) < 4 {
		return [2]int{-1, -1}
	}

	diff := 0

	for _, num := range nums {
		diff ^= num
	}

	res := [2]int{diff, diff}
	last1Bit := (diff & (diff - 1)) ^ diff
	for _, num := range nums {
		if num&last1Bit == 0 {
			res[0] ^= num
		} else {
			res[1] ^= num
		}
	}

	return res
}
