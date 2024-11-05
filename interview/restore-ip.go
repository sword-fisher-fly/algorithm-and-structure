package interview

import (
	"strconv"
	"strings"
)

// 1) 25525511135 -> ["255.255.11.135", "255.255.111.35"]
// 2) 0000 -> ["0.0.0.0"]
// 3) 110110 -> ["1.10.11.0", "1.10.1.10", "1.1.0.110", "11.0.1.10", "11.0.11.0", "110.1.1.0"]
// 4) "010010" -> ["0.10.0.10", "0.100.1.0"]

func RestoreIpAddresses(s string) []string {
	isValidIpPart := func(s string, i, j int) bool {
		if s[i] == '0' && i != j {
			return false
		}

		num, err := strconv.Atoi(s[i : j+1])
		if err != nil {
			return false
		}

		return num >= 0 && num <= 255
	}

	res := []string{}
	path := []string{}

	var dfs func(int)

	dfs = func(start int) {
		if start == len(s) && len(path) == 4 {
			res = append(res, strings.Join(path, "."))
			return
		}

		// 剪枝
		if start >= len(s) || len(path) >= 4 {
			return
		}

		for i := start; i < len(s); i++ {
			if isValidIpPart(s, start, i) {
				path = append(path, s[start:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)

	return res
}

func RestoreIpAddressesII(s string) []string {
	isDigit := func(c byte) bool {
		return c >= '0' && c <= '9'
	}

	isValidIpPart := func(s string, i, j int) bool {
		if s[i] == '0' && i != j {
			return false
		}

		num := 0
		for i < j+1 {
			if !isDigit(s[i]) {
				return false
			}

			num = num*10 + int(s[i]-'0')
			i++
		}

		return num >= 0 && num <= 255
	}

	res := []string{}
	path := []string{}

	var dfs func(int)

	dfs = func(start int) {
		if start == len(s) && len(path) == 4 {
			res = append(res, strings.Join(path, "."))
			return
		}

		// 剪枝
		if start >= len(s) || len(path) >= 4 {
			return
		}

		for i := start; i < len(s) && i < start+3; i++ {
			if isValidIpPart(s, start, i) {
				path = append(path, s[start:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)

	return res
}
