package interview

// 表达式求值

// https://www.cnblogs.com/irvingluo/p/13301157.html
// https://www.nowcoder.com/practice/c215ba61c8b1443b996351df929dc4d4?tpId=308&tqId=1076787&ru=%2Fexam%2Foj&qru=%2Fta%2Falgorithm-start%2Fquestion-ranking&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AC%2594%25E8%25AF%2595%26topicId%3D308&dayCountBigMember=365%E5%A4%A9

// 递归写法: 一层层去除括号
func EvalExpr(s string) int {
	num := 0
	sign := byte('+')
	stack := []int{}

	isDigit := func(c byte) bool {
		return c >= '0' && c <= '9'
	}

	isOperator := func(c byte) bool {
		return c == '+' || c == '-' || c == '*' || c == '/' || c == '%'
	}

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}

		if isDigit(s[i]) {
			num = num*10 + int(s[i]-'0')
		}

		// 巧妙配对, 去除最外层()
		if s[i] == '(' {
			j := i + 1
			counter := 1
			for counter > 0 {
				if s[j] == '(' {
					counter++
				}
				if s[j] == ')' {
					counter--
				}
				j++
			}

			// s[j-1] = ')'
			// fmt.Printf("s[i:j]=%s, s[i+1][j-1]=%s\n", s[i:j], s[i+1:j-1])
			num = EvalExpr(s[i+1 : j-1])
			i = j - 1 // for i++
		}

		if isOperator(s[i]) || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top*num)
			case '/':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top/num)
			case '%':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top%num)
			}

			sign = s[i]
			num = 0
		}
	}

	ans := 0

	for i := range stack {
		ans += stack[i]
	}

	return ans
}

// TODO:
// 中缀表达式
// 后缀表达式
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/
// https://zq99299.github.io/dsalg-tutorial/dsalg-java-hsp/05/05.html#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0
