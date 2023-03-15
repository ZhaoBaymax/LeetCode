package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	cur_max := 0
	var dfs func(r *TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		left, right := dfs(r.Left), dfs(r.Right)
		cur_left := 0
		cur_right := 0
		if r.Left != nil && r.Val == r.Left.Val {
			cur_left = left + 1
		}
		if r.Right != nil && r.Val == r.Right.Val {
			cur_right = right + 1
		}
		cur_max = max(cur_max, cur_right+cur_left)
		return max(cur_left, cur_right)
	}
	dfs(root)
	return cur_max
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
