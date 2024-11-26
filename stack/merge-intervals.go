package stack

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param intervals Interval类一维数组
 * @return Interval类一维数组
 */
type intervals []*Interval

func (l intervals) Len() int {
	return len(l)
}

func (l intervals) Less(i, j int) bool {
	if l[i].Start == l[j].Start && l[i].End < l[j].End {
		return true
	}

	return l[i].Start < l[j].Start
}

func (l intervals) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (i *Interval) String() string {
	return fmt.Sprintf("[%d, %d]", i.Start, i.End)
}

func printIntervals(arr []*Interval) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("i=%d [%d, %d]\n", i, arr[i].Start, arr[i].End)
	}
}

func MergeIntervals(arr []*Interval) []*Interval {
	// write code here
	if len(arr) <= 1 {
		return arr
	}

	sort.Sort(intervals(arr))
	stack := []*Interval{arr[0]}
	for i := 1; i < len(arr); i++ {
		top := stack[len(stack)-1]
		if top.End < arr[i].Start {
			stack = append(stack, arr[i])
		} else {
			if top.End < arr[i].End {
				top.End = arr[i].End
			}
		}
	}

	return stack
}
