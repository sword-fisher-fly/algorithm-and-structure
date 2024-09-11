package graph

type Node struct {
	ID    int64
	Edges []*Edge
}

type Edge struct {
	FromNode *Node
	EndNode  *Node
}

type Graph struct {
	Nodes map[int64]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int64]*Node),
	}
}

func (g *Graph) AddNode(id int64) *Node {
	if _, exist := g.Nodes[id]; !exist {
		g.Nodes[id] = &Node{
			ID:    id,
			Edges: make([]*Edge, 0),
		}
	}

	return g.Nodes[id]
}

func (g *Graph) AddEdge(fromNode, endNode *Node, direction bool) {
	if fromNode == nil || endNode == nil {
		return
	}

	edge := &Edge{
		FromNode: fromNode,
		EndNode:  endNode,
	}

	revEdge := &Edge{
		FromNode: endNode,
		EndNode:  fromNode,
	}

	if direction {
		fromNode.Edges = append(fromNode.Edges, edge)
	} else {
		fromNode.Edges = append(fromNode.Edges, edge)
		endNode.Edges = append(endNode.Edges, revEdge)
	}
}
