package str

import "fmt"

// input: 1+3-2*4+4
//
//	 +1+3-2*4+4
//		 +-1+3-2*4+4
//
// 巧妙点
// 1) 头部添加'+', 扫描数字在前而导致左操作数先被push到stack中
// 2) '-'表示为负数, 因而stack内所有元素可看作待相加求和的数
func Calculate(s string) int {
	n := len(s)

	x := 0
	var sign byte = '+'

	stack := []int{}

	for i := 0; i < n; i++ {
		if isDigit(s[i]) {
			x = x*10 + int(s[i]-'0')
		}

		if i == n-1 || !isDigit(s[i]) && s[i] != ' ' {
			fmt.Printf("stack= %v, x=%d, sign=%c, s[%d]=%c\n", stack, x, sign, i, s[i])
			switch sign {
			case '+':
				stack = append(stack, x)
			case '-':
				stack = append(stack, -x)
			case '*':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top*x)
			case '/':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top/x)
			}

			x = 0
			sign = s[i]
		}
	}

	sum := 0
	for i := range stack {
		sum += stack[i]
	}

	return sum
}
