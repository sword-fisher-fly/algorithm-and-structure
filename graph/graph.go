package graph

func TraversalGraphByDFS(g *Graph, x int64, n int64, result *[][]int64, path []int64) {
	if x == n {
		*result = append(*result, append(append([]int64{}, path...), n))
		return
	}

	if v, exist := g.Nodes[x]; exist {
		for _, edge := range v.Edges {
			path = append(path, x)
			TraversalGraphByDFS(g, edge.EndNode.ID, n, result, path)
			path = path[:len(path)-1]
		}
	}
}
