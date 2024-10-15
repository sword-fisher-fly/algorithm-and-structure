package offer

import (
	"container/heap"
	"sort"
)

// heap.Interface
// type Interface interface {
// 	sort.Interface
// 	Push(x any) // add x as element Len()
// 	Pop() any   // remove and return element Len() - 1.
// }

type KthLargest struct {
	sort.IntSlice
	k int
}

// TODO: 有点困惑??, 有些问题
func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]

	return v
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)

	if kl.Len() > kl.k {
		heap.Pop(kl)
	}

	return kl.IntSlice[0]
}

func (kl *KthLargest) Top() int {
	return kl.IntSlice[0]
}
