package huawei

func CountSameChars(s string, t byte) int {
	isSameAlpha := func(c byte, t byte) bool {

		// fmt.Printf("c=%c, t=%c\n", c, t)
		if c >= 'a' && c <= 'z' {
			return c == t || c-'a'+'A' == t
		}

		if c >= 'A' && c <= 'Z' {
			return c == t || c-'A'+'a' == t
		}

		return c == t
	}

	res := 0

	for i := 0; i < len(s); i++ {
		if isSameAlpha(s[i], t) {
			res++
		}
	}

	return res
}
