package stack

type MaxSlideWindowQueue struct {
	Queue []int
}

func NewMaxSlideWindowQueue() *MaxSlideWindowQueue {
	return &MaxSlideWindowQueue{
		Queue: make([]int, 0),
	}
}

func (s *MaxSlideWindowQueue) Push(x int) {
	for len(s.Queue) > 0 && x >= s.Queue[len(s.Queue)-1] {
		s.Queue = s.Queue[:len(s.Queue)-1]
	}

	s.Queue = append(s.Queue, x)
}

func (s *MaxSlideWindowQueue) Pop(x int) {
	if len(s.Queue) > 0 && x == s.Queue[0] {
		s.Queue = s.Queue[1:]
	}
}

// 滑动窗口包含k个元素, 依次从左至右在数组上滑动, 求滑动窗口里的最大值
// 
func MaxInSlideWindow(arr []int, k int) []int {
	if len(arr) < k {
		return []int{}
	}

	queue := NewMaxSlideWindowQueue()
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
