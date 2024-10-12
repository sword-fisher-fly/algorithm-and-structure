package offer

func IsPopOrder(pushV, popV []int) bool {
	if len(pushV) != len(popV) {
		return false
	}

	n := len(pushV)

	j := 0

	s := []int{}
	for i := 0; i < len(pushV); i++ {
		for j < n && (len(s) == 0 || s[len(s)-1] != popV[i]) {
			s = append(s, pushV[j])
			j++
		}

		if popV[i] == s[len(s)-1] {
			s = s[:len(s)-1]
			continue
		}

		return false
	}

	return true
}
