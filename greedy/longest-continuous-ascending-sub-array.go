package greedy

import "fmt"

func LongestContinuousAscendingSubSequence(arr []int) int {
	cnt := 1
	result := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			cnt++
		} else {
			cnt = 1
		}

		fmt.Printf("cnt=%d, arr[%d]=%d\n", cnt, i, arr[i])

		result = max(result, cnt)
	}

	return result
}
