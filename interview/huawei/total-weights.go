package huawei

func TotalDifferentWeights(weights []int, nums []int) int {
	w, n := len(weights), len(nums)
	if w != n {
		return -1
	}

	vv := []int{}

	for i := 0; i < w; i++ {
		for j := 0; j < nums[i]; j++ {
			vv = append(vv, weights[i])
		}
	}

	cloneMap := func(m map[int]struct{}) map[int]struct{} {
		t := make(map[int]struct{})
		for k := range m {
			t[k] = struct{}{}
		}

		return t
	}

	s := make(map[int]struct{})
	s[0] = struct{}{}

	for i := 0; i < len(vv); i++ {
		t := cloneMap(s)
		for j := range t {
			s[j+vv[i]] = struct{}{}
		}
	}

	return len(s)
}
