package stack

type Stack []int64

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() int64 {
	if s.Empty() {
		return -1
	}

	val := (*s)[s.size()-1]
	*s = (*s)[:s.size()-1]
	return val
}

func (s *Stack) Top() int64 {
	if s.Empty() {
		return -1
	}

	return (*s)[s.size()-1]
}

func (s *Stack) Push(v int64) {
	*s = append(*s, v)
}

func (s *Stack) size() int {
	return len(*s)
}

func (s *Stack) Empty() bool {
	return s.size() == 0
}
