package main

import (
	"fmt"
)

type Graph struct {
	nodes []*GraphNode
}

type GraphNode struct {
	id    int
	edges map[int]int
}

var visited = map[int]bool{}

func New() *Graph {
	return &Graph{
		nodes: []*GraphNode{},
	}
}

func (g *Graph) AddNode() (id int) {
	id = len(g.nodes)
	g.nodes = append(g.nodes, &GraphNode{
		id:    id,
		edges: make(map[int]int),
	})
	return
}

func (g *Graph) AddEdge(n1, n2 int, w int) {
	g.nodes[n1].edges[n2] = w
}

func (g *Graph) DFSUtil(node *GraphNode) {
	visited[node.id] = true
	for edge := range node.edges {
		if !visited[edge] {
			for _, node := range g.nodes {
				if node.id == edge {
					g.DFSUtil(node)
				}
			}
		}
	}
}

func (g *Graph) DFS() map[int]bool {
	for _, node := range g.nodes {
		if !visited[node.id] {
			g.DFSUtil(node)
		}
	}
	return visited
}

func main() {
	g := new(Graph)
	g.AddNode()
	g.AddNode()
	g.AddNode()
	g.AddNode()
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 0, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 3, 1)

	fmt.Print(g.DFS())
}
