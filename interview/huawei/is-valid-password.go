package huawei

import "strings"

// https://www.nowcoder.com/practice/184edec193864f0985ad2684fbc86841?tpId=37&tqId=21243&rp=1&ru=/exam/oj/ta&qru=/exam/oj/ta&sourceUrl=%2Fexam%2Foj%2Fta%3FtpId%3D37&difficulty=undefined&judgeStatus=undefined&tags=&title=
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isOther(c byte) bool {
	if c == ' ' || c == '\n' {
		return false
	}

	return !isDigit(c) && !isLower(c) && !isUpper(c)
}

func IsValidPassword(s string) bool {
	if len(s) <= 8 {
		return false
	}
	// dArr[0]: digit
	// dArr[1]: lower
	// dArr[2]: upper
	// dArr[3]: other
	// dArr := [4]bool{}
	// atLeastThree := func(a string) bool {
	// 	if len(a) == 0 {
	// 		return false
	// 	}

	// 	for i := 0; i < len(s); i++ {
	// 		if isDigit(s[i]) {
	// 			dArr[0] = true
	// 			continue
	// 		}
	// 		if isLower(s[i]) {
	// 			dArr[1] = true
	// 			continue
	// 		}
	// 		if isUpper(s[i]) {
	// 			dArr[2] = true
	// 			continue
	// 		}

	// 		if !isOther(s[i]) {
	// 			return false
	// 		}
	// 		dArr[3] = true
	// 	}

	// 	count := 0
	// 	for i := range dArr {
	// 		if dArr[i] {
	// 			count++
	// 		}
	// 	}

	// 	return count >= 3
	// }

	atLeastThree := func(s string) bool {
		a, b, c, d := 0, 0, 0, 0
		for i := 0; i < len(s); i++ {
			if isDigit(s[i]) {
				a = 1
			} else if isLower(s[i]) {
				b = 1
			} else if isUpper(s[i]) {
				c = 1
			} else {
				if !isOther(s[i]) {
					return false
				}
				d = 1
			}
		}
		return a+b+c+d >= 3
	}

	if !atLeastThree(s) {
		return false
	}

	hasSameSubStr := func(s string) bool {
		for i := 0; i < len(s)-2; i++ {
			ss := strings.Split(s, s[i:i+3])
			if len(ss) >= 3 {
				return false
			}
		}
		return true
	}

	if !hasSameSubStr(s) {
		return false
	}

	return true
}
