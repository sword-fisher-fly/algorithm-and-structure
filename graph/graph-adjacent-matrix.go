package graph

import (
	"fmt"
	"strings"
)

// 1 -> 3
// 3 -> 5
// 1 -> 2
// 2 -> 4
// 4 -> 5
type GraphMatrix struct {
	n          int       // the count of node
	nodeMatrix [][]int64 // i,j represent node, nodeMatrix[i][j] !=0 means there exists an edge, its value represent the weight of edge
}

func NewGraphMatrix(n int64) *GraphMatrix {
	g := &GraphMatrix{
		n:          int(n),
		nodeMatrix: make([][]int64, n+1),
	}

	for i, _ := range g.nodeMatrix {
		g.nodeMatrix[i] = make([]int64, n+1)
	}

	return g
}

func (g *GraphMatrix) AddEdge(startPoint, endPoint int64) {
	if startPoint < 1 || startPoint > int64(g.n) || endPoint < 1 || endPoint > int64(g.n) {
		return
	}

	g.nodeMatrix[startPoint][endPoint] = 1
}

func (g *GraphMatrix) DFS(x, n int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	path = append(path, x)

	minLen := 0

	var dfs func(x, n int)
	dfs = func(x, n int) {
		if x == n {
			tmp := make([]int, len(path))
			copy(tmp, path)
			if len(tmp) < minLen || minLen == 0 {
				minLen = len(tmp)
			}
			res = append(res, tmp)

			return
		}

		for i := 1; i <= n; i++ {
			if g.nodeMatrix[x][i] == 1 {
				path = append(path, i)
				dfs(i, n)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(x, n)

	fmt.Printf("minLen: %d\n", minLen)

	return res
}

func (g *GraphMatrix) PrettyPrintGraphMatrix() string {
	res := strings.Builder{}
	for i := 1; i < len(g.nodeMatrix); i++ {
		for j := 1; j < len(g.nodeMatrix[0]); j++ {
			if g.nodeMatrix[i][j] == 1 {
				res.WriteString(fmt.Sprintf("%d->%d\n", i, j))
			}
		}
	}

	return res.String()
}
