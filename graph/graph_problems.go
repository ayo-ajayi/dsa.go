package graph

// find out if there's a route between two nodes in a directed graph
func IsRouteBetweenNodes(node1 *Node, node2 *Node) bool {

	for _, v := range node1.DepthFirstSearch() {
		if v == node2.data {
			return true
		}
	}

	for _, v := range node2.DepthFirstSearch() {
		if v == node1.data {
			return true
		}
	}

	return false
}

// go through a DFS and terminate early before completing search
func IsRouteBetweenTwoNodes(node1 *Node, node2 *Node) bool {
	return node1.DFSearch(node2) || node2.DFSearch(node1)
}

func (n *Node) DFSearch(target *Node) bool {
	visited := make(map[*Node]bool)
	return n.dfsHelper(target, visited)
}

func (n *Node) dfsHelper(target *Node, visited map[*Node]bool) bool {
	if n == target {
		return true
	}
	visited[n] = true
	for _, adjacent := range n.adjacent {
		if !visited[adjacent] && adjacent.dfsHelper(target, visited) {
			return true
		}
	}
	return false
}
