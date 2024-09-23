package str

// 1) -0      true
// 2) +0       true
// 3) -012     false
// 4) +021     false
// 5) 0123     false
// 6) 1.       false
// 7) 1.E      false
// 8) 1.0E-2   true
// 9) 1..0     false
// 10) 1.EE2   false
// 11) 1.E-02  false
// 12) 1.0E-02 false
// 13) 1.01E+03 false
// 14) 1.12E+3  true
// 15) 12       true
// 16) +123     true
// 17) -189     true
// 18) +1.      false
// 19) +1.01E+3 true

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isNonZeroInterger(s string) bool {
	hasSign := false

	if len(s) < 1 {
		return false
	}

	if s[0] == '+' || s[0] == '-' {
		hasSign = true
	}

	if hasSign && len(s) == 1 {
		return false
	}

	if hasSign {
		if s[1] == '0' {
			return false
		}

		for i := 1; i < len(s); i++ {
			if !isDigit(s[i]) {
				return false
			}
		}
	} else {
		if s[0] == '0' {
			return false
		}

		for i := 0; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				return false
			}
		}
	}

	return true
}

func isValidIntergerAfterDot(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := range s {
		if !isDigit(s[i]) {
			return false
		}
	}

	return true
}

func isValidInterger(s string) bool {
	hasSign := false

	if len(s) == 0 {
		return false
	}

	if s[0] == '+' || s[0] == '-' {
		hasSign = true
	}

	if hasSign && len(s) == 1 {
		return false
	}

	if hasSign {
		if s[1] == '0' && len(s) > 2 {
			return false
		}

		for i := 1; i < len(s); i++ {
			if !isDigit(s[i]) {
				return false
			}
		}
	} else {
		if s[0] == '0' && len(s) > 1 {
			return false
		}

		for i := 0; i < len(s); i++ {
			if !isDigit(s[i]) {
				return false
			}
		}
	}

	return true
}
