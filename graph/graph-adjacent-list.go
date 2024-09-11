package graph

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
