package offer

import "container/heap"

// 返回数组中相同元素出现频次前k个元素值

// 最小堆思想

type pair struct{ v, cnt int }
type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].cnt < h[j].cnt
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}

func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]

	return v
}

func TopKFrequence(nums []int, k int) []int {
	numCntMap := make(map[int]int)
	for i := range nums {
		numCntMap[nums[i]]++
	}

	h := hp{}
	for v, cnt := range numCntMap {
		heap.Push(&h, pair{v, cnt})
		if len(h) > k {
			heap.Pop(&h)
		}
	}

	ans := []int{}

	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(&h).(pair).v)
	}

	return ans
}

func TopKFrequenceII(nums []int, k int) []int {
	maxCount := 0

	numCntMap := make(map[int]int)
	for i := range nums {
		numCntMap[nums[i]]++
		if numCntMap[nums[i]] > maxCount {
			maxCount = numCntMap[nums[i]]
		}
	}

	ans := []int{}
	for k > 0 {
		for v, cnt := range numCntMap {
			if cnt == maxCount {
				ans = append(ans, v)
				delete(numCntMap, v)
				k--
				// note: if it has found the k top items, then break the for-loop
				if k == 0 {
					break
				}
			}
		}

		maxCount--
	}

	return ans
}
