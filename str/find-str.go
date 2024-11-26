package str

import (
	"regexp"
)

func FindStr(s, t string) int {
	re := regexp.MustCompile(t)

	idx := re.FindStringIndex(s)

	if idx == nil {
		return -1
	}

	return idx[0]
}

func FindStrII(s, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return -1
	}

	start := -1
	var k int
	for i := 0; i <= m-n; i++ {
		start = i
		for j := i; j < n && j < m; j++ {
			if t[j] == '[' {
				for k = j; k < n && t[k] != ']'; {
					k++
				}
			} else if s[i] == t[j] {
				k++
				j++
			} else {
				break
			}
		}
	}

	return start
}
