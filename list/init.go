package list

var (
	sortedArray1 = []int{1, 2, 4, 7, 8, 9, 14}
	sortedArray2 = []int{11, 20, 43, 56, 64, 78, 82, 94}
	sortedList1  *ListNode
	sortedList2  *ListNode

	randomArray1               = []int{8, 4, 5, 10, 7, 9, 14}
	expectedSortedRandomArray1 = []int{4, 5, 7, 8, 9, 10, 14}
	randomArray2               = []int{101, 300, 45, 2, 3, 56, 87, 94}
	expectedSortedRandomArray2 = []int{2, 3, 45, 56, 87, 94, 101, 300}
	randomList1                *ListNode
	randomList2                *ListNode

	expectedMergedSortedListArray []int
)

func init() {
	sortedList1 = NewListFromArray(sortedArray1)
	sortedList2 = NewListFromArray(sortedArray2)

	expectedMergedSortedListArray = mergeSortedArray(sortedArray1, sortedArray2)

	randomList1 = NewListFromArray(randomArray1)
	randomList2 = NewListFromArray(randomArray2)
}

func mergeSortedArray(arr1, arr2 []int) []int {
	arr := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}

	if i < len(arr1) {
		arr = append(arr, arr1[i:]...)
	}

	if j < len(arr2) {
		arr = append(arr, arr2[j:]...)
	}

	return arr
}
