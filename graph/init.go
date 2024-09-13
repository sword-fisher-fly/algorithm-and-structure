package graph

var (
	vertex5Vertex4Edge = [][2]int64{
		{0, 1},
		{0, 2},
		{1, 3},
		{2, 4},
	}
)

var directions [4][2]int

var (
	graphList5Vertex4Edge *GraphList
)

func init() {
	graphList5Vertex4Edge = NewGraphList()
	graphList5Vertex4Edge.BuildGraphList(vertex5Vertex4Edge)

	// (x, y) 右 -> 下 -> 左 -> 上
	directions = [4][2]int{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}
}
