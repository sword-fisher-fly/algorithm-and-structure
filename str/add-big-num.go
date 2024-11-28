package str

func AddBigNum(s1, s2 string) string {
	m, n := len(s1), len(s2)

	if m > n {
		s1, s2 = s2, s1
		m, n = n, m
	}

	res := make([]byte, 0, n+1)
	carry := 0
	for i, j := m-1, n-1; i >= 0; i, j = i-1, j-1 {
		a1 := int(s1[i] - '0')
		a2 := int(s2[j] - '0')

		sum := a1 + a2 + carry

		res = append(res, byte(sum%10)+'0')
		carry = sum / 10
	}

	for i := n - m - 1; i >= 0; i-- {
		sum := int(s2[i]-'0') + carry
		res = append(res, byte(sum%10)+'0')
		carry = sum / 10
	}

	if carry > 0 {
		res = append(res, byte(carry)+'0')
	}

	lr := len(res)
	for i := 0; i < lr/2; i++ {
		res[i], res[lr-i-1] = res[lr-i-1], res[i]
	}

	return string(res)
}
