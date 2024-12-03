package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func preorder(tn *TreeNode, result *[]int) {
	if tn == nil {
		return
	}
	*result = append(*result, tn.val)
	preorder(tn.left, result)
	preorder(tn.right, result)

}

func inorder(tn *TreeNode, result *[]int) {
	if tn == nil {
		return
	}
	preorder(tn.left, result)
	*result = append(*result, tn.val)
	preorder(tn.right, result)
}

func postorder(tn *TreeNode, result *[]int) {
	if tn == nil {
		return
	}
	preorder(tn.right, result)
	preorder(tn.left, result)
	*result = append(*result, tn.val)
}

func main() {

	root := &TreeNode{val: 6}

	root.left = &TreeNode{val: 9}
	root.right = &TreeNode{val: 5}
	root.left.left = &TreeNode{val: 4}
	root.left.right = &TreeNode{val: 6}
	root.right.left = &TreeNode{val: 1}
	root.right.right = &TreeNode{val: 2}

	res := make([]int, 0)
	preorder(root, &res)
	fmt.Println(res)

	res = make([]int, 0)
	inorder(root, &res)
	fmt.Println(res)

	res = make([]int, 0)
	postorder(root, &res)
	fmt.Println(res)

}
