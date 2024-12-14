package offer

import (
	"testing"
)

func TestKthLargest(t *testing.T) {
	kl := &KthLargest{
		IntSlice: []int{3, 1, 2, 5, 4},
		k:        10,
	}

	kl.Add(6)
	kl.Add(7)
	kl.Add(8)
	kl.Add(9)
	kl.Add(10)
	kl.Add(11)
	t.Log(kl.Top())

	kl.Add(12)

	t.Log(kl.Top())

	kl.Add(13)
	t.Log(kl.Top())

	kl.Add(7)
	t.Log(kl.Top())

	kl.Add(8)
	t.Log(kl.Top())

	t.Logf("Heap size: %d", kl.Len())

	for i := 0; i < 10; i++ {
		item := kl.Pop()
		t.Logf("Pop item %d -> %d", i, item)
	}
}
