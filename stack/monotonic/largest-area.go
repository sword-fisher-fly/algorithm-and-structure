package monotonic

import "fmt"

// 1) 暴力解法
// 2) 双指针
// 3）单调栈
func LargestRectangleAreaByForce(heights []int) int {
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		left := i
		right := i
		for left > 0 && heights[left-1] >= heights[i] {
			left--
		}
		for right < len(heights)-1 && heights[right+1] >= heights[i] {
			right++
		}
		area := (right - left + 1) * heights[i]

		maxArea = max(area, maxArea)
	}

	return maxArea
}

func LargestRectangleAreaByTwoPointer(heights []int) int {
	if len(heights) < 1 {
		return 0
	}

	minLeftIndex := make([]int, len(heights))
	minRightIndex := make([]int, len(heights))

	minLeftIndex[0] = -1
	for i := 1; i < len(heights); i++ {
		t := i - 1
		for t >= 0 && heights[t] >= heights[i] {
			t = minLeftIndex[t]
		}

		minLeftIndex[i] = t
	}

	minRightIndex[len(heights)-1] = len(heights)
	for i := len(heights) - 2; i >= 0; i-- {
		t := i + 1
		for t < len(heights) && heights[t] >= heights[i] {
			t = minRightIndex[t]
		}

		minRightIndex[i] = t
	}

	sum := 0
	for i := 0; i < len(heights); i++ {
		area := (minRightIndex[i] - minLeftIndex[i] - 1) * heights[i]
		sum = max(area, sum)
	}

	return sum
}

func LargestRectangleArea(heights []int) int {
	newHeights := make([]int, len(heights)+2)

	copy(newHeights[1:], heights)

	fmt.Printf("newHeight=%v\n", newHeights)

	stack := []int{0}
	maxArea := 0
	for i := 1; i < len(newHeights); i++ {
		for newHeights[i] < newHeights[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			w := i - stack[len(stack)-1] - 1
			h := newHeights[mid]

			maxArea = max(maxArea, w*h)
		}

		stack = append(stack, i)
	}

	return maxArea
}
