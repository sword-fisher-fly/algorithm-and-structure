package graph

import "testing"

func graphInitialize() *Graph {
	g := NewGraph()

	node1 := g.AddNode(1)
	node2 := g.AddNode(2)
	node3 := g.AddNode(3)
	node4 := g.AddNode(4)

	// 1 -> 2 -> 3 -> 4
	// 1 -> 3 -> 4
	// 1 -> 2 -> 4

	g.AddEdge(node1, node2, true)
	g.AddEdge(node2, node3, true)
	g.AddEdge(node3, node4, true)

	g.AddEdge(node1, node3, true)

	g.AddEdge(node2, node4, true)

	return g
}

func TestGraphInitialize(t *testing.T) {
	g := graphInitialize()

	for id, node := range g.Nodes {
		for _, edge := range node.Edges {
			t.Logf("node: %v, edge: %v-%v", id, edge.FromNode.ID, edge.EndNode.ID)
		}
	}
}

func TestGraphDFS(t *testing.T) {
	g := graphInitialize()

	var result [][]int64
	var path []int64

	dfs(g, 1, 4, &result, path)

	for i, path := range result {
		t.Logf("path %v: %v", i, path)
	}
}
