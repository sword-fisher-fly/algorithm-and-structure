package stack

import "strconv"

func EvalRPN(tokens []string) int64 {
	stack := []int64{}

	for _, token := range tokens {
		i, err := strconv.ParseInt(token, 10, 64)
		if err == nil {
			stack = append(stack, i)
		} else {
			leftNum := stack[len(stack)-2]
			rightNum := stack[len(stack)-1]

			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, leftNum+rightNum)
			case "-":
				stack = append(stack, leftNum-rightNum)
			case "*":
				stack = append(stack, leftNum*rightNum)
			case "/":
				stack = append(stack, leftNum/rightNum)
			}
		}
	}

	return stack[0]
}
