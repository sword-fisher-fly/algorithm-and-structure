package special

import (
	"container/heap"
	"sort"
)

func GetLeastNumbers(input []int, k int) []int {
	// write code here
	if k == 0 || k > len(input) {
		return []int{}
	}

	sort.Ints(input)

	return input[:k]
}

func GetLeastNumbersII(input []int, k int) []int {
	if k == 0 || k > len(input) {
		return []int{}
	}

	var quickSort func([]int, int, int)

	partion := func(arr []int, low, high int) int {
		p := arr[high]
		i := low - 1
		for j := low; j < high; j++ {
			if arr[j] < p {
				i++ // 指向大于p的值
				if i != j {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}

		arr[i+1], arr[high] = arr[high], arr[i+1]

		return i + 1
	}
	quickSort = func(arr []int, low int, high int) {
		if low >= high {
			return
		}

		pivot := partion(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}

	quickSort(input, 0, len(input)-1)

	return input[:k]
}

type pq []int

func (h pq) Len() int {
	return len(h)
}

func (h pq) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h pq) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pq) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *pq) Pop() any {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return top
}

func GetLeastNumbersIII(input []int, k int) []int {
	if k == 0 || k > len(input) {
		return []int{}
	}

	h := &pq{}
	heap.Init(h)
	for i := range input {
		heap.Push(h, input[i])
		if h.Len() > k {
			top := heap.Pop(h).(int)
			if input[i] < top {
				heap.Push(h, input[i])
			} else {
				heap.Push(h, top)
			}
		}
	}
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(int))
	}

	return res
}
