package main

import (
	"fmt"
)

type Queue struct {
	s []int
}

func (q *Queue) Push(x int) {
	q.s = append(q.s, x)
}

func (q *Queue) Pop() int {
	if q.IsEmpty() {
		return -1
	}

	val := q.s[0]
	q.s = q.s[1:]
	return val
}

func (q Queue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.s[0]
}

func (q Queue) IsEmpty() bool {
	return len(q.s) == 0
}

func main() {
	n := 0

	fmt.Scan(&n)

	q := &Queue{s: make([]int, 0)}

	for i := 0; i < n; i++ {
		var op string
		fmt.Scan(&op)
		switch op {
		case "push":
			num := 0
			fmt.Scan(&num)
			q.Push(num)
		case "pop":
			if q.IsEmpty() {
				fmt.Println("error")
				continue
			}

			fmt.Println(q.Pop())
		case "front":
			if q.IsEmpty() {
				fmt.Println("error")
				continue
			}
			fmt.Println(q.Front())
		}
	}
}
