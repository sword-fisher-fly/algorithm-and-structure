package sort

func partition(arr []int, low, high int) {
	if low+1 >= high {
		return
	}
	first, last := low, high-1
	key := arr[last]
	for first < last {
		// 左半部分
		for first < last && arr[first] <= key {
			first++
		}
		arr[last] = arr[first]

		// 右半部分
		for first < last && arr[last] >= key {
			last--
		}
		arr[first] = arr[last]
	}

	arr[first] = key
	partition(arr, low, first)
	partition(arr, first+1, high)
}
func QuickSort(arr []int) []int {
	partition(arr, 0, len(arr))
	return arr
}
