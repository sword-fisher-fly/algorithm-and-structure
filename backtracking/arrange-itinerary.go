package backtracking

// TODO:
// 给定一个机票字符串二维数组[from, to]
// from: 飞机出发点
// to:   飞机降落点
// 假定所有机票至少存在一个合理行程
// 所有机票必须都用一次且只用一次
// targets: map[string]map[string]int
// from -> to: 出发地 -> 目的地
func FindItinerary(startSite string, tickets [][2]string) []string {
	result := []string{startSite}
	targets := make(map[string]map[string]int)
	for _, t := range tickets {
		from, to := t[0], t[1]
		// fmt.Printf("From: %s, To: %s\n", from, to)
		if _, ok := targets[from]; !ok {
			targets[from] = make(map[string]int)
		}
		targets[from][to]++
	}

	// fmt.Printf("targets: %v\n", targets)

	var backtracking func(tickNums int) bool
	backtracking = func(tickNums int) bool {
		if len(result) == tickNums+1 {
			return true
		}

		lastSite := result[len(result)-1]
		for to, num := range targets[lastSite] {
			if num > 0 {
				result = append(result, to)
				targets[lastSite][to]--
				if backtracking(tickNums) {
					return true
				}
				result = result[:len(result)-1]
				targets[lastSite][to]++
			}
		}

		return false
	}

	backtracking(len(tickets))

	return result
}
