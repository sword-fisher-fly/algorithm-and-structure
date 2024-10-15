package offer

import (
	"fmt"
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
	fmt.Printf("Add 10 kth largest: %v\n", kl.IntSlice)

	kl.Add(11)
	fmt.Printf("Add 11 kth largest: %v\n", kl.IntSlice)

	t.Log(kl.Top())

	kl.Add(12)
	fmt.Printf("Add 12 kth largest: %v\n", kl.IntSlice)

	t.Log(kl.Top())

	kl.Add(13)
	fmt.Printf("Add 13 kth largest: %v\n", kl.IntSlice)

	t.Log(kl.Top())

	kl.Add(7)
	fmt.Printf("kth largest: %v\n", kl.IntSlice)
	t.Log(kl.Top())

	kl.Add(8)
	fmt.Printf("kth largest: %v\n", kl.IntSlice)
	t.Log(kl.Top())
}
