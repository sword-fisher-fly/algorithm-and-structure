package stack

func IsValidBracket(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	bracketPair := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}

	for i := range s {
		switch s[i] {
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if top != bracketPair[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]

		default:
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
