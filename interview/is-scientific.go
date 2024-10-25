package interview

//https://leetcode.cn/problems/valid-number/solutions/564188/you-xiao-shu-zi-by-leetcode-solution-298l/

// 下面的都是有效数字："2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"，而接下来的不是："abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"
func isRealNumber(s string, mustNum bool) bool {
	if len(s) == 0 {
		return false
	}

	hasDot, hasNum := false, false
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	isDigit := func(c byte) bool {
		return c >= '0' && c <= '9'
	}

	for i := range s {
		if isDigit(s[i]) {
			hasNum = true
		} else if s[i] == '.' {
			if hasDot || mustNum {
				return false
			}

			hasDot = true
		} else {
			return false
		}
	}

	return hasNum
}

func isNumber(s string) bool {
	idx := -1
	for i := range s {
		switch s[i] {
		case 'e', 'E':
			idx = i
			break
		}
	}

	if idx == 0 || idx == len(s)-1 {
		return false
	}
	// 可能是科学记数法
	if idx != -1 {
		return isRealNumber(s[:idx], false) && isRealNumber(s[idx+1:], true)
	}

	// 不可能是科学记数法
	return isRealNumber(s, false)
}
