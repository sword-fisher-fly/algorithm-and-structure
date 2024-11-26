package huawei

//https://www.nowcoder.com/practice/96bd6684e04a44eb80e6a68efc0ec6c5?tpId=295&tqId=23260&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3F...%3D
const MOD = 1e9 + 7

func mergeSort(arr []int, low, high int, count *int) {
	if low < high {
		mid := low + (high-low)/2
		mergeSort(arr, low, mid, count)
		mergeSort(arr, mid+1, high, count)
		merge(arr, low, mid, high, count)
	}
}

func merge(arr []int, low, mid, high int, count *int) {
	temp := make([]int, high-low+1)
	idx := 0

	sIdx, l := low, low
	r := mid + 1

	for l <= mid && r <= high {
		if arr[l] <= arr[r] {
			temp[idx] = arr[l]
			idx++
			l++
		} else {
			temp[idx] = arr[r]
			*count = *count + mid + 1 - l
			*count = (*count) % MOD
			idx++
			r++
		}
	}

	for l <= mid {
		temp[idx] = arr[l]
		idx++
		l++
	}

	for r <= high {
		temp[idx] = arr[r]
		idx++
		r++
	}

	for i := range temp {
		arr[sIdx] = temp[i]
		sIdx++
	}
}

func InversePairs(nums []int) int {
	// write code here
	if len(nums) < 2 {
		return 0
	}
	count := 0
	mergeSort(nums, 0, len(nums)-1, &count)

	return count
}
