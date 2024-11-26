package sort

import "container/heap"

// 前K个高频元素

type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func TopKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)

	for _, v := range nums {
		numMap[v]++
	}
	h := &IHeap{}

	heap.Init(h)

	// min-Heap
	for key, val := range numMap {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}

	return res
}
