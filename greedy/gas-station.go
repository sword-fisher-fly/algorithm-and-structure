package greedy

func CanCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 || len(cost) == 0 {
		return -1
	}

	total := 0
	current := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		current += gas[i] - cost[i]

		if current < 0 {
			current = 0
			start = i + 1
		}
	}

	if total < 0 {
		return -1
	}

	return start
}

// 暴力求解
func CanCompleteCircuitII(gas []int, cost []int) int {
	if len(gas) == 0 || len(cost) == 0 {
		return -1
	}

	var index int
	for i := 0; i < len(gas); i++ {
		rest := gas[i] - cost[i]
		index = (index + 1) % len(gas)
		for rest > 0 && index != i {
			rest += gas[index] - cost[index]
			index = (index + 1) % len(gas)
		}

		if rest >= 0 && index == i {
			return index
		}
	}

	return -1
}
