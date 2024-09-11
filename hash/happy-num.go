package hash

import "fmt"

// getSum return the sum of  the squares of the numbers at each position of it
func getSum(n int64) int64 {
	var ret int64

	for n != 0 {
		ret += (n % 10) * (n % 10)
		n /= 10
	}

	return ret
}

func dumpMapKeys(m map[int64]struct{}) []int64 {
	var ret []int64

	for k := range m {
		ret = append(ret, k)
	}

	return ret
}

func IsHappy(n int64) bool {
	m := make(map[int64]struct{})

	for n != 1 {
		n = getSum(n)

		if _, ok := m[n]; ok {
			fmt.Printf("keys: %v\n", dumpMapKeys(m))

			return false
		}

		m[n] = struct{}{}
	}

	return true
}
