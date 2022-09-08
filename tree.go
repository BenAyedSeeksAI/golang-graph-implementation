package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func createNode(value int) *Node {
	node := &Node{}
	node.value = value
	node.left = nil
	node.right = nil
	return node
}
func mainTree() {
	root := createNode(10)
	root.left = createNode(9)
	root.left.left = createNode(7)
	root.left.right = createNode(8)
}
