package backtracking

import (
	"strings"
)

// [start, end]
func isValidNum(s string, start, end int) bool {
	if start > end {
		return false
	}

	if end-start > 2 {
		return false
	}

	if s[start] == '0' && start != end {
		return false
	}

	num := 0
	for i := start; i <= end; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}

		num = num*10 + int(s[i]-'0')

		if num > 255 {
			return false
		}
	}

	return true
}

func restoreIPAddressHelper(s string, startIndex int, pointNum int, path []string, result *[]string) {
	if startIndex == len(s) && pointNum == 4 {
		*result = append(*result, strings.Join(path, "."))
		return
	}

	for i := startIndex; i < len(s); i++ {
		if !isValidNum(s, startIndex, i) {
			break // ok too??
			// continue // ok
		}

		path = append(path, s[startIndex:i+1])
		restoreIPAddressHelper(s, i+1, pointNum+1, path, result)
		path = path[:len(path)-1]
	}
}
func RestoreIPAddresses(s string) []string {
	result := []string{}
	restoreIPAddressHelper(s, 0, 0, []string{}, &result)

	return result
}
