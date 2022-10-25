package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	return traversal(root, &res)
}

func traversal(root *TreeNode, res *[]int) []int {
	if root == nil {
		return nil
	}
	*res = append(*res, root.Val)
	traversal(root.Left, res)
	traversal(root.Right, res)
	return *res
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(preorderTraversal(root))
}
