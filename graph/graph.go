package graph

import (
	"github.com/ayo-ajayi/dsa.go/queue"
)

type Node struct {
	data     any
	adjacent []*Node
}

func NewNode(data any) *Node {
	return &Node{
		data: data,
	}
}

// set node both ways for adjacency matrix and set node once with adjacency list
func (n *Node) SetAdjacent(nodes []*Node) {
	n.adjacent = nodes
}

func (root *Node) DepthFirstSearch() []any {
	results := make([]any, 0)
	visited := map[*Node]bool{}
	root.depthFirstSearch(&results, visited)
	return results
}

func (root *Node) BreadthFirstSearch() []any {
	results := make([]any, 0)
	visited := make(map[*Node]bool)
	root.breadthFirstSearch(&results, visited)
	return results
}

func (root *Node) depthFirstSearch(results *[]any, visited map[*Node]bool) {
	if root == nil {
		return
	}
	*results = append(*results, root.data)
	visited[root] = true

	for _, r := range root.adjacent {
		if !visited[r] {
			r.depthFirstSearch(results, visited)
		}
	}
}

func (root *Node) breadthFirstSearch(results *[]any, visited map[*Node]bool) {
	q := queue.NewQueue()
	visited[root] = true
	q.Add(root)

	for !q.IsEmpty() {
		r := q.Remove().(*Node)
		*results = append(*results, r.data)

		for _, node := range r.adjacent {
			if !visited[node] {
				visited[node] = true
				q.Add(node)
			}
		}
	}
}
