package offer

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	stack := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		top := stack[len(stack)-1]
		if top[1] >= intervals[i][0] {
			if top[1] < intervals[i][1] {
				stack[len(stack)-1] = []int{top[0], intervals[i][1]}
			}
		} else {
			stack = append(stack, intervals[i])
		}
	}

	return stack
}
