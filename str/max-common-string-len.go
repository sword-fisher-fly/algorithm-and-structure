package str

import "strings"

func FindMaxCommonSubStringLength(s1, s2 string) int {
	m, n := len(s1), len(s2)

	if m > n {
		s1, s2 = s2, s1
		m, n = n, m
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}
	maxVal := 0

	for i := 0; i < m; i++ {
		for j := m; j > i; j-- {
			if strings.Contains(s2, s1[i:j]) {
				maxVal = max(maxVal, j-i)
			}
		}
	}

	return maxVal
}
