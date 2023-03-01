package main

func rob(root *TreeNode) int {
	mem := make(map[*TreeNode]int, 0)
	return robRange(root, mem)
}

func robRange(root *TreeNode, mem map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if _, ok := mem[root]; ok {
		return mem[root]
	}
	l, r := 0, 0
	if root.Left != nil {
		l = robRange(root.Left.Left, mem) + robRange(root.Left.Right, mem)
	}
	if root.Right != nil {
		r = robRange(root.Right.Left, mem) + robRange(root.Right.Right, mem)
	}
	get_it := root.Val + l + r
	not_get := robRange(root.Left, mem) + robRange(root.Right, mem)
	res := max(get_it, not_get)
	mem[root] = res
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
