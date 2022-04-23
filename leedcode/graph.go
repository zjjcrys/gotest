package leedcode

type Node struct {
	Val       int
	Neighbors []*Node
}

//topic 133
func cloneGraph(node *Node) *Node {
	visited := map[*Node]*Node{} //使用make 效果一样
	var cg func(node *Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if _, ok := visited[node]; ok {
			return visited[node]
		}

		cloneNode := &Node{node.Val, []*Node{}}
		visited[node] = cloneNode
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(n))
		}

		return cloneNode
	}

	return cg(node)

}
