package lapuladong

func FindDuplicateNum(nums []int) int {
	n := len(nums)
	l, r := 0, n-1

	ans := 0
	for l <= r {
		mid := l + (r-l)>>1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}

		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}

	return ans
}
