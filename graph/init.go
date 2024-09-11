package graph

var (
	vertex5Vertex4Edge = [][2]int64{
		{0, 1},
		{0, 2},
		{1, 3},
		{2, 4},
	}
)

var (
	graphList5Vertex4Edge *GraphList
)

func init() {
	graphList5Vertex4Edge = NewGraphList()
	graphList5Vertex4Edge.BuildGraphList(vertex5Vertex4Edge)
}
