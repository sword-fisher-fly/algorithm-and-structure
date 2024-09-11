package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	t.Logf("stack: %v", *s)

	val5 := s.Pop()
	if val5 != 5 {
		t.Errorf("pop value is not 5")
	}

	val4 := s.Pop()
	if val4 != 4 {
		t.Errorf("pop value is not 4")
	}

	val3 := s.Pop()
	if val3 != 3 {
		t.Errorf("pop value is not 3")
	}

	if s.Top() != 2 {
		t.Errorf("top value is not 2")
	}

	val2 := s.Pop()
	if val2 != 2 {
		t.Errorf("pop value is not 2")
	}

	val1 := s.Pop()
	if val1 != 1 {
		t.Errorf("pop value is not 1")
	}

	if !s.Empty() {
		t.Errorf("stack is not empty")
	}
}
