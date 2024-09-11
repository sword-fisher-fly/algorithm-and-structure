package graph

type GraphMatrix struct {
	n          int       // the count of node
	nodeMatrix [][]int64 // i,j represent node, nodeMatrix[i][j] !=0 means there exists an edge, its value represent the weight of edge
}

func NewGraphMatrix(n int64) *GraphMatrix {
	g := &GraphMatrix{
		n:          int(n),
		nodeMatrix: make([][]int64, n),
	}

	for i, _ := range g.nodeMatrix {
		g.nodeMatrix[i] = make([]int64, n)
	}

	return g
}

func (g *GraphMatrix) AddEdge(startPoint, endPoint int64) {
	if startPoint < 0 || startPoint >= int64(g.n) || endPoint < 0 || endPoint >= int64(g.n) {
		return
	}

	g.nodeMatrix[startPoint][endPoint] = 1
}
