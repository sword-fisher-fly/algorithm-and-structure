package sort

func BubbleSort(arr []int) []int {

	//input: 3,6,8,10,4,2,1,5,7,9
	//steps:
	// 1) 1,3,6,8,10,4,2,5,7,9
	// 2) 1,2,3,6,8,10,4,5,7,9
	// 3) 1,2,3,4,6,8,10,5,7,9
	// 4) 1,2,3,4,5,6,8,10,7,9
	// 5) 1,2,3,4,5,6,7,8,10,9
	// 6) 1,2,3,4,5,6,7,8,9,10
	for i := 0; i < len(arr); i++ {
		swapped := false
		for j := len(arr) - 1; j > i; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}
