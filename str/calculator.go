package str

import "fmt"

type Priority int

const (
	NoPriority = iota
	AddMinusPriority
	MultiplyDividePriority
)

func isOperator(c byte) bool {
	switch c {
	case '-', '+', '*', '/', '%':
		return true
	}

	return false
}

func getPriority(op string) int {
	switch op {
	case "*", "/", "%":
		return MultiplyDividePriority
	case "+", "-":
		return AddMinusPriority
	}

	return NoPriority
}

func calculate(numLeft, numRight int, op string) int {
	switch op {
	case "+":
		return numLeft + numRight
	case "-":
		return numLeft - numRight
	case "*":
		return numLeft * numRight
	case "/":
		return numLeft / numRight
	case "%":
		return numLeft % numRight
	}

	return 0
}

func EvalExpr(input string) int {
	operandStack := make([]int, 0)
	operatorStack := make([]string, 0)

	var popAndCalculateFromOperandStackAndOperatorStack func()
	popAndCalculateFromOperandStackAndOperatorStack = func() {
		leftNum, rightNum := operandStack[len(operandStack)-2], operandStack[len(operandStack)-1]
		operandStack = operandStack[:len(operandStack)-2]

		op := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]

		operandStack = append(operandStack, calculate(leftNum, rightNum, op))
	}

	var addNewOperator func(string)
	addNewOperator = func(newOp string) {
		for len(operatorStack) > 0 && getPriority(operatorStack[len(operatorStack)-1]) >= getPriority(newOp) {
			popAndCalculateFromOperandStackAndOperatorStack()
		}

		operatorStack = append(operatorStack, newOp)
	}

	var prev byte // 上一个属于运算符的字符

	for i := 0; i < len(input); i++ {
		// prev是否为一个运算符
		if !isOperator(input[i]) && prev > 0 {
			newOp := string(prev)
			addNewOperator(newOp)
			prev = 0
		}

		if isDigit(input[i]) { // refactor here to use for-loop
			curVal := int(input[i] - '0')
			if i > 0 && isDigit(input[i-1]) {
				topVal := operandStack[len(operandStack)-1]
				topVal = topVal*10 + curVal

				operandStack = operandStack[:len(operandStack)-1]
				operandStack = append(operandStack, topVal)
			} else {
				operandStack = append(operandStack, curVal)
			}
		}

		if isOperator(input[i]) {
			if prev > 0 {
				// TODO 处理两个字符运算符
				prev = 0
			} else {
				prev = input[i]
			}
		}

		if input[i] == '(' {
			operatorStack = append(operatorStack, "(")
		}

		if input[i] == ')' {
			for operatorStack[len(operatorStack)-1] != "(" {
				popAndCalculateFromOperandStackAndOperatorStack()
			}

			// Don't forget to pop the '('
			operatorStack = operatorStack[:len(operatorStack)-1]
		}
	}

	fmt.Printf("len(operatorStack)=%d, len(operandStack)=%d\n", len(operatorStack), len(operandStack))
	for len(operatorStack) > 0 {
		popAndCalculateFromOperandStackAndOperatorStack()
	}

	return operandStack[0]
}

// Optimize for more efficient.
func EvalExprII(input string) int {
	operandStack := make([]int, 0)
	operatorStack := make([]byte, 0)

	var calculate func(int, int, byte) int
	calculate = func(leftNum, rightNum int, op byte) int {
		switch op {
		case '+':
			return leftNum + rightNum
		case '-':
			return leftNum - rightNum
		case '*':
			return leftNum * rightNum
		case '/':
			return leftNum / rightNum
		case '%':
			return leftNum % rightNum
		}

		return 0
	}

	var popAndCalculateFromOperandStackAndOperatorStack func()
	popAndCalculateFromOperandStackAndOperatorStack = func() {
		leftNum, rightNum := operandStack[len(operandStack)-2], operandStack[len(operandStack)-1]
		operandStack = operandStack[:len(operandStack)-2]

		op := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]

		operandStack = append(operandStack, calculate(leftNum, rightNum, op))
	}

	var getPriority func(byte) int

	getPriority = func(op byte) int {
		switch op {
		case '*', '/', '%':
			return MultiplyDividePriority
		case '+', '-':
			return AddMinusPriority
		}

		return NoPriority
	}

	var addNewOperator func(byte)
	addNewOperator = func(newOp byte) {
		for len(operatorStack) > 0 && getPriority(operatorStack[len(operatorStack)-1]) >= getPriority(newOp) {
			popAndCalculateFromOperandStackAndOperatorStack()
		}

		operatorStack = append(operatorStack, newOp)
	}

	for i := 0; i < len(input); {
		if isDigit(input[i]) { // refactor here to use for-loop
			curVal := int(input[i] - '0')

			for i = i + 1; i < len(input) && isDigit(input[i]); {
				curVal = curVal*10 + int(input[i]-'0')
			}

			operandStack = append(operandStack, curVal)
			continue
		}

		if isOperator(input[i]) {
			addNewOperator(input[i])
		}

		if input[i] == '(' {
			operatorStack = append(operatorStack, '(')
		}

		if input[i] == ')' {
			for operatorStack[len(operatorStack)-1] != '(' {
				popAndCalculateFromOperandStackAndOperatorStack()
			}

			// Don't forget to pop the '('
			operatorStack = operatorStack[:len(operatorStack)-1]
		}

		i++
	}

	for len(operatorStack) > 0 {
		popAndCalculateFromOperandStackAndOperatorStack()
	}

	return operandStack[0]
}
