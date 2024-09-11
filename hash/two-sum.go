package hash

func TwoSum(nums []int64, target int64) [2]int {
	ret := [2]int{-1, -1}

	// key: val -> num: idx
	numMap := make(map[int64]int)

	// Example: 2,7,11,15 target: 9
	for i, num := range nums {
		if idx, ok := numMap[target-num]; ok {
			ret[0] = idx
			ret[1] = i
			break
		}

		numMap[num] = i
	}

	return ret
}
