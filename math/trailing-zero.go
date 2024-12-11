package math

func TrailingZeroes(n int) int {
	res := 0
	for n > 0 {
		n = n / 5
		res += n
	}

	return res
}

func TrailingZerosII(n int) int {
	res := 0
	for i := 5; i <= n; i += 5 {
		temp := i
		for temp%5 == 0 {
			temp = temp / 5
			res++
		}
	}

	return res
}

func TrailingZeroesIII(n int) int {
	T
	if n < 5 {
		return 0
	}

	if n == 5 {
		return 1
	}

	return n/5 + TrailingZeroesIII(n/5)
}
