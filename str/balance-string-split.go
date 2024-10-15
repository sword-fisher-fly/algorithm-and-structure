package str

func BalanceStringSplit(s string) int {
	balance := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			balance++
		} else {
			balance--
		}

		if balance == 0 {
			count++
		}
	}

	return count
}
