package backtracking

func PalindromePartition(s string) [][]string {
	res := [][]string{}
	path := []string{}

	// 左闭合右闭合区间
	isPalindrome := func(s string, start, end int) bool {
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	var backtracking func(s string, startIndex int, path []string)
	backtracking = func(s string, startIndex int, path []string) {
		if startIndex >= len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := startIndex; i < len(s); i++ {
			if !isPalindrome(s, startIndex, i) {
				continue
			}

			path = append(path, s[startIndex:i+1])
			backtracking(s, i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtracking(s, 0, path)

	return res
}

func computePalindrome(s string) [][]bool {
	res := make([][]bool, len(s))
	for i := range res {
		res[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				res[i][j] = true
			} else if j-i == 1 {
				res[i][j] = s[i] == s[j]
			} else {
				res[i][j] = (s[i] == s[j]) && res[i+1][j-1]
			}
		}
	}

	return res
}

func PalindromePartitionII(s string) [][]string {
	res := [][]string{}
	path := []string{}

	// 左闭合右闭合区间
	isPalindrome := computePalindrome(s)

	var backtracking func(s string, startIndex int, path []string)
	backtracking = func(s string, startIndex int, path []string) {
		if startIndex >= len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := startIndex; i < len(s); i++ {
			if !isPalindrome[startIndex][i] {
				continue
			}

			path = append(path, s[startIndex:i+1])
			backtracking(s, i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtracking(s, 0, path)

	return res
}
