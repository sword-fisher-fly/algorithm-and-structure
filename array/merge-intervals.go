package array

import "sort"

type Points [][2]int

func (p Points) Len() int {
	return len(p)
}
func (p Points) Less(i, j int) bool {
	if p[i][0] < p[j][0] {
		return true
	} else if p[i][0] == p[j][0] && p[i][1] < p[j][1] {
		return true
	} else {
		return false
	}
}
func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func MergeIntervals(intervals [][2]int) [][2]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Sort(Points(intervals))

	stack := [][2]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		top := stack[len(stack)-1] // [1,3] [2,4]
		if intervals[i][0] <= top[1] {
			stack[len(stack)-1] = [2]int{top[0], max(top[1], intervals[i][1])}
		} else {
			stack = append(stack, intervals[i])
		}
	}

	return stack
}
