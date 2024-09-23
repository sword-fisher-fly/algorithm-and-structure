package str

import (
	"fmt"
	"strings"
)

// 1)  0
// 2) +0
// 3) -0
// 4) 1.0
// 5) -1.0
// 6) +1.0
func isValidRealNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	hasDot := false

	fields := strings.Split(s, ".")
	if len(fields) == 2 {
		hasDot = true
	}

	if hasDot {
		fmt.Printf("integer part: %s, decimal part: %s\n", fields[0], fields[1])
		return isValidInterger(fields[0]) && isValidIntergerAfterDot(fields[1])
	}

	return isValidInterger(s)
}
