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

// 3, 2, 1, 4, 5, 10, 8, 7, 6, 9
// low: 0 high: 9
func partitionII(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	// i+1: 指向第一个大于等于pivot的元素
	// j: 指向第一个小于pivot的元素
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			// fmt.Printf("i: %d, j: %d, arr[i]: %d, arr[j]: %d\n", i, j, arr[i], arr[j])
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

// 左闭右闭区间
// 3, 2, 1, 4, 5, 10, 8, 7, 6, 9
func QuickSortII(arr []int) {
	var quickSort func([]int, int, int)

	quickSort = func(arr []int, low, high int) {
		if low < high {
			pivotIndex := partitionII(arr, low, high)
			quickSort(arr, low, pivotIndex-1)
			quickSort(arr, pivotIndex+1, high)
		}
	}

	quickSort(arr, 0, len(arr)-1)
}
