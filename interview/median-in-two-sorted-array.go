package interview

// 总数为偶数
// nums1: 1,5,7,8,9
// nums2: 2,3,4,6,10
// return: (5+6)/2
// 总数为奇数
// nums1: 1,5,7,8,9
// nums2: 2,3,4,6,10,11
// return: 6

// 1,5,/7,8,9 i:= 2
// 2,3,4,/6,10 j:= 3

// 1,5,/7,8,9
// 2,3,4,6,/10
func FindMedianInTwoSortedArrays(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	if m > n {
		return FindMedianInTwoSortedArrays(nums2, nums1)
	}

	iMin := 0
	iMax := m
	// [iMin, iMax)
	// m+n =7
	// i+j = (7+1)/2 = 4  0,1,2,3,4,5,6
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (m+n+1)/2 - i

		if i > 0 && j != n && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else if i != m && j > 0 && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return maxLeft
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}

			return (maxLeft + minRight) / 2
		}
	}

	return -1
}
