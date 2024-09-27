package greedy

func LongestContinuousAscendingSubSequence(arr []int) int {
	cnt := 1
	result := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			cnt++
		} else {
			cnt = 1
		}

		result = max(result, cnt)
	}

	return result
}
