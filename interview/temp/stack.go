package main

import (
	"fmt"
	"math"
)

type Stack []int

func NewStack() *Stack {
	ret := &Stack{}
	*ret = make([]int, 0)

	return ret
}

func (s Stack) Len() int {
	return len(s)
}

func (s *Stack) Pop() {
	if s.Len() > 0 {
		(*s) = (*s)[:len(*s)-1]
	} else {
		fmt.Println("error")
	}
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s Stack) Top() int {
	if s.Len() > 0 {
		return s[s.Len()-1]
	} else {
		fmt.Println("error!")
	}

	return math.MaxInt64
}

func main() {
	s := NewStack()
	var n int
	fmt.Scan(&n)

	fmt.Printf("n: %d\n", n)

	for i := 0; i < n; i++ {
		var op string
		fmt.Scan(&op)

		switch op {
		case "push":
			var b int
			fmt.Scan(&b)
			s.Push(b)
		case "pop":
			s.Pop()
		case "top":
			fmt.Println(s.Top())
		}
	}
}
