package main

import "fmt"

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
func treeTraversal(root *Node) {
	if root != nil {
		if root.left != nil {
			treeTraversal(root.left)
		}
		fmt.Println(root.value)
		if root.right != nil {
			treeTraversal(root.right)
		}
	}

}
func mainTree() {
	root := createNode(10)
	root.left = createNode(9)
	root.left.left = createNode(7)
	root.left.right = createNode(8)
	treeTraversal(root)
}
