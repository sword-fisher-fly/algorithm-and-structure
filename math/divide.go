package math

import "math"

// -2^31 <= a, b <= 2^31 - 1
func Divide(a, b int) int {
	if b == 1 {
		return a
	}

	if a == -math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	sign := false
	if a > 0 && b > 0 || a < 0 && b < 0 {
		sign = true
	}

	if a > 0 {
		a = -a
	}

	if b > 0 {
		b = -b
	}

	ans := 0

	for a <= b {
		x := b
		cnt := 1
		for x >= math.MinInt32>>1 && a <= x+x {
			x += x
			cnt += cnt
		}

		a -= x
		ans += cnt
	}

	if sign {
		return ans
	}

	return -ans
}
