package stack

type MinSlideWindowQueue struct {
	Queue []int
}

func NewMinSlideWindowQueue() *MinSlideWindowQueue {
	return &MinSlideWindowQueue{Queue: []int{}}
}

// queue: from low to high
// 4, 3, 5, 4, 3, 3, 6, 7
func (q *MinSlideWindowQueue) Push(x int) {
	for len(q.Queue) > 0 && x < q.Queue[len(q.Queue)-1] {
		q.Queue = q.Queue[:len(q.Queue)-1]
	}

	q.Queue = append(q.Queue, x)
}

func (q *MinSlideWindowQueue) Pop(x int) {
	if len(q.Queue) > 0 && x == q.Queue[0] {
		q.Queue = q.Queue[1:]
	}
}

func MinInSlideWindow(arr []int, k int) []int {
	if len(arr) < k {
		return []int{}
	}

	queue := NewMinSlideWindowQueue()

	for i := 0; i < k; i++ {
		queue.Push(arr[i])
	}

	res := []int{queue.Queue[0]}
	for i := k; i < len(arr); i++ {
		queue.Pop(arr[i-k])
		queue.Push(arr[i])
		res = append(res, queue.Queue[0])
	}

	return res
}
