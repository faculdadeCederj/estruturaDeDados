package main

type Node struct {
	Left  *Node
	Value int
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (n *Node) addLeftNode(node *Node) {
	n.Left = node
}

func (n *Node) addRightNode(node *Node) {
	n.Right = node
}
