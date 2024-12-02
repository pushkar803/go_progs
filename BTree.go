package main

import (
	"fmt"
)

type bTreeNode struct {
	data  int
	left  *bTreeNode
	right *bTreeNode
}

func main10() {

	root := &bTreeNode{}
	root = insertBtree(root, 5)
	root = insertBtree(root, 10)
	root = insertBtree(root, 6)
	x := searchbTree(root, 6)
	fmt.Println(x)

}

func insertBtree(node *bTreeNode, data int) *bTreeNode {
	if node == nil {
		return &bTreeNode{data: data}
	}
	if data < node.data {
		node.left = insertBtree(node.left, data)
	} else {
		node.right = insertBtree(node.right, data)
	}

	return node
}

func searchbTree(node *bTreeNode, data int) bool {
	if node == nil {
		return false
	}

	if data == node.data {
		return true
	}

	if data < node.data {
		return searchbTree(node.left, data)
	}

	return searchbTree(node.right, data)

}
