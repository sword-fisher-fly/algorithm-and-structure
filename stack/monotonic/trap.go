package monotonic

// N根柱子, nums数组表示其高度
// 求N根柱子最大续接雨水

// 1) 暴力求解
// 2）双指针
// 3）单调栈

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func TrapByForce(heights []int) int {
	if len(heights) < 3 {
		return 0
	}

	sum := 0

	for i := 1; i < len(heights)-1; i++ {
		lHeight := heights[i]
		rHeight := heights[i]

		for j := 0; j < i; j++ {
			if heights[j] > lHeight {
				lHeight = heights[j]
			}
		}

		for j := i + 1; j < len(heights); j++ {
			if heights[j] > rHeight {
				rHeight = heights[j]
			}
		}

		h := min(lHeight, rHeight) - heights[i]

		if h > 0 {
			sum += h
		}
	}

	return sum
}

func TrapByTwoPointer(heights []int) int {
	if len(heights) < 3 {
		return 0
	}

	maxLeft := make([]int, len(heights))
	maxRight := make([]int, len(heights))
	maxLeft[0] = heights[0]
	maxRight[len(heights)-1] = heights[len(heights)-1]

	for i := 1; i < len(heights); i++ {
		maxLeft[i] = max(maxLeft[i-1], heights[i])
	}

	for j := len(heights) - 2; j >= 0; j-- {
		maxRight[j] = max(maxRight[j+1], heights[j])
	}

	sum := 0

	for i := 1; i < len(heights)-1; i++ {
		h := min(maxLeft[i], maxRight[i]) - heights[i]
		if h > 0 {
			sum += h
		}
	}

	return sum
}
