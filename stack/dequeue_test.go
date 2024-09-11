package stack

import "testing"

func TestDequeue(t *testing.T) {
	q := NewDequeue()

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)

	val1 := q.Pop()
	if val1 != 1 {
		t.Errorf("Expected 1, got %d", val1)
	}

	val2 := q.Pop()
	if val2 != 2 {
		t.Errorf("Expected 2, got %d", val2)
	}

	val3 := q.Pop()
	if val3 != 3 {
		t.Errorf("Expected 3, got %d", val3)
	}

	val4 := q.Pop()
	if val4 != 4 {
		t.Errorf("Expected 4, got %d", val4)
	}

	val5 := q.Pop()
	if val5 != 5 {
		t.Errorf("Expected 5, got %d", val5)
	}

	if !q.Empty() {
		t.Errorf("Expected empty queue, got %v", q)
	}
}
