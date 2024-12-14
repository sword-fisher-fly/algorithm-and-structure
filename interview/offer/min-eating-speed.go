package offer

func MinEatingSpeed(piles []int, h int) int {
	if len(piles) < 1 {
		return 0
	}
	// count = 30; piles[i]=31 or 60 h=1
	// t = piles[i] + count -1 /count
	maxVal := 0
	for i := range piles {
		maxVal = max(maxVal, piles[i])
	}

	sum := func(count int) int {
		cost := 0
		for i := range piles {
			cost += (piles[i] + count - 1) / count
		}

		return cost
	}

	left, right := 1, maxVal
	for left < right {
		mid := left + (right-left)>>1
		s := sum(mid)
		if s <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
