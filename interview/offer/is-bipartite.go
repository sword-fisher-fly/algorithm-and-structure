package offer

// https://leetcode.cn/problems/is-graph-bipartite/
func IsBipartite(graph [][]int) bool {
	var find func(x int, p []int) int

	find = func(x int, p []int) int {
		if p[x] != x {
			p[x] = find(p[x], p)
		}

		return p[x]
	}

	n := len(graph)

	p := make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = i
	}

	for u := 0; u < n; u++ {
		for _, v := range graph[u] {
			if find(u, p) == find(v, p) {
				return false
			}

			p[find(v, p)] = find(graph[u][0], p)
		}
	}

	return true
}
