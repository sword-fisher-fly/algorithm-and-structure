package graph

import (
	"fmt"
	"strings"
)

type GraphList struct {
	Vertex map[int64]struct{}
	Edges  map[int64][]int64
}

func NewGraphList() *GraphList {
	return &GraphList{
		Vertex: make(map[int64]struct{}),
		Edges:  make(map[int64][]int64),
	}
}

func (g *GraphList) AddEdge(fromNode, endNode int64) {
	g.Vertex[fromNode] = struct{}{}
	g.Vertex[endNode] = struct{}{}
	g.Edges[fromNode] = append(g.Edges[fromNode], endNode)
}

func (g *GraphList) BuildGraphList(edge [][2]int64) {
	for _, e := range edge {
		g.AddEdge(e[0], e[1])
	}
}

func (g *GraphList) GetVertex() []int64 {
	var vertex []int64
	for k := range g.Vertex {
		vertex = append(vertex, k)
	}

	return vertex
}

func (g *GraphList) DFS(x, n int64) [][]int64 {
	res := make([][]int64, 0)
	path := make([]int64, 0)
	path = append(path, x)

	minLen := 0
	var dfs func(x, n int64)
	dfs = func(x, n int64) {
		if x == n {
			tmp := make([]int64, len(path))
			copy(tmp, path)
			if len(tmp) < minLen || minLen == 0 {
				minLen = len(tmp)
			}

			res = append(res, tmp)
			return
		}

		for _, v := range g.Edges[x] {
			path = append(path, v)
			dfs(v, n)
			path = path[:len(path)-1]
		}
	}

	dfs(x, n)

	fmt.Printf("minLen: %d\n", minLen)

	return res
}

func (g *GraphList) PrettyPrintGraphList() string {
	res := strings.Builder{}

	for _, v := range g.GetVertex() {
		// res.WriteString(fmt.Sprintf("%d -> ", v))
		for _, e := range g.Edges[v] {
			res.WriteString(fmt.Sprintf("%d->%d\n", v, e))
		}
	}

	return res.String()
}
