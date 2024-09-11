package str

import "strings"

// 1)  0
// 2) +0
// 3) -0
// 4) 1.0
// 5) -1.0
// 6) +1.0
func isValidRealNumber(s string) bool {
	if len(s) < 1 {
		return false
	}

	hasDot := false

	fields := strings.Split(s, ".")
	if len(fields) == 2 {
		hasDot = true
	}

	if hasDot {
		return isValidInterger(fields[0]) && isValidIntergerAfterDot(fields[1])
	}

	return isValidInterger(s)
}
