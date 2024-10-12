package main

import (
	"fmt"
)

type CircleQueue struct {
	s     []int
	limit int
}

func (c CircleQueue) IsEmpty() bool {
	return len(c.s) == 0
}

func (c CircleQueue) Len() int {
	return len(c.s)
}

func (c CircleQueue) IsFull() bool {
	return len(c.s) == c.limit
}

func (c *CircleQueue) Push(x int) {
	if c.IsFull() {
		return
	}

	c.s = append(c.s, x)
}

func (c *CircleQueue) Pop() int {
	if c.IsEmpty() {
		return -1
	}

	val := c.s[0]
	c.s = c.s[1:]

	return val
}

func (c CircleQueue) Front() int {
	if c.IsEmpty() {
		return -1
	}

	return c.s[0]
}

// func main() {
//     // n:循环队列容量 p:操作次数
//     var n, p int
//     fmt.Scan(&n, &p)
//     cq := &CircleQueue{limit: n}
//     scanner := bufio.NewScanner(os.Stdin)

//     step := 0
//     var op string
//     var num int
//     for scanner.Scan() && step < n {
//         input := scanner.Text()
//         fields := strings.Fields(input)
//         op = fields[0]
//         switch op {
//         case "push":
//             if cq.IsFull() {
//                 fmt.Println("full")
//                 continue
//             }
//             num, _ = strconv.Atoi(fields[1])
//             cq.Push(num)
//         case "pop":
//             if cq.IsEmpty() {
//                 fmt.Println("empty")
//                 continue
//             }
//             fmt.Println(cq.Pop())
//         case "front":
//             if cq.IsEmpty() {
//                 fmt.Println("empty")
//                 continue
//             }
//             fmt.Println(cq.Front())
//         }
//     }

//     if err := scanner.Err(); err != nil {
//         fmt.Println(err)
//     }

// }

func main() {
	// n:循环队列容量 p:操作次数
	var n, p int
	fmt.Scan(&n, &p)
	cq := &CircleQueue{limit: n}

	var op string
	var num int
	for i := 0; i < p; i++ {
		fmt.Scanln(&op, &num)
		switch op {
		case "push":
			if cq.IsFull() {
				fmt.Println("full")
				continue
			}
			cq.Push(num)
		case "pop":
			if cq.IsEmpty() {
				fmt.Println("empty")
				continue
			}
			fmt.Println(cq.Pop())
		case "front":
			if cq.IsEmpty() {
				fmt.Println("empty")
				continue
			}
			fmt.Println(cq.Front())
		}
	}
}
