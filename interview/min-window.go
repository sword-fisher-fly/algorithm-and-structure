package interview

import (
	"fmt"
	"math"
)

// 1) 统计s和t出现的字符个数
// 2) 左右指针移动
func MinWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	sCntMap := map[byte]int{}
	tCntMap := map[byte]int{}

	for i := range t {
		tCntMap[t[i]]++
		sCntMap[s[i]]++
	}

	// s, t
	check := func() bool {
		for k, v := range tCntMap {
			if sCntMap[k] < v {
				return false
			}
		}

		fmt.Println("==== check ====")
		return true
	}

	left, right := -1, -1
	if check() {
		return s[0:len(t)]
	}

	minLen := math.MaxInt

	l, r := 0, len(t)

	fmt.Printf("l=%d, r=%d\n", l, r)

	for r < len(s) {
		if tCntMap[s[r]] > 0 {
			sCntMap[s[r]]++
		}

		for check() && l <= r {
			if r-l+1 < minLen {
				minLen = r - l + 1
				left = l
				right = l + minLen
			}

			if tCntMap[s[l]] > 0 && sCntMap[s[l]] > 0 {
				sCntMap[s[l]]--
			}

			l++
		}

		r++
	}

	if minLen != math.MaxInt {
		return s[left:right]
	}

	return ""
}
