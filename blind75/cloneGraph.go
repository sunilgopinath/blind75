package blind75

/**
 * Definition for a Node.

 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return cloneHelper(node, visited)
}

func cloneHelper(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := visited[node]; ok {
		return v
	}
	visited[node] = &Node{Val: node.Val}
	for _, v := range node.Neighbors {
		visited[node].Neighbors = append(visited[node].Neighbors, cloneHelper(v, visited))
	}
	return visited[node]
}
