package main

import "fmt"

func main() {
	root := Node{Value: 1}

	binaryTree := BinaryTree{Root: &root}

	root.addLeftNode(&Node{Value: 3})
	root.addRightNode(&Node{Value: 4})

	root.Left.addLeftNode(&Node{Value: 2})

	root.Right.addLeftNode(&Node{Value: 5})
	root.Right.addRightNode(&Node{Value: 3})

	fmt.Println(binaryTree.Root.Right.Right)
}
