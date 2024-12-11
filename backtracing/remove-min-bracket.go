package backtracking

//https://leetcode.cn/problems/remove-invalid-parentheses/?envType=problem-list-v2&envId=2cktkvj
func RemoveInvalidParentheses(s string) []string {
	lremove, rremove := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			lremove++
		} else if s[i] == ')' {
			if lremove == 0 {
				rremove++
			} else {
				lremove--
			}
		}
	}

	isValidBracket := func(s string) bool {
		cnt := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				cnt++
			} else if s[i] == ')' {
				cnt--
				if cnt < 0 {
					return false
				}
			}
		}
		return cnt == 0
	}

	res := []string{}
	var backtracing func(string, int, int, int)

	backtracing = func(s string, start int, lremove int, rremove int) {
		if lremove == 0 && rremove == 0 {
			if isValidBracket(s) {
				res = append(res, s)
			}
			return
		}

		for i := start; i < len(s); i++ {
			// 去重
			if i > start && s[i] == s[i-1] {
				continue
			}

			// if lremove+rremove > len(s)-i {
			//     return
			// }

			if lremove > 0 && s[i] == '(' {
				backtracing(s[:i]+s[i+1:], i, lremove-1, rremove)
			}

			if rremove > 0 && s[i] == ')' {
				backtracing(s[:i]+s[i+1:], i, lremove, rremove-1)
			}
		}
	}

	backtracing(s, 0, lremove, rremove)

	return res
}
