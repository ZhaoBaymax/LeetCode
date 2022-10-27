package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BSTSequences(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return [][]int{}
	}
	q := list.New()
	q.PushBack(root)
	var dfs func(q *list.List, track []int)
	dfs = func(q *list.List, track []int) {
		if q.Len() == 0 {
			temp := make([]int, len(track))
			copy(temp, track)
			ans = append(ans, temp)
			return
		}

		for i := 0; i < q.Len(); i++ {
			cur := q.Remove(q.Front()).(*TreeNode)
			track = append(track, cur.Val)
			if cur.Left != nil {
				q.PushBack(cur.Left)
			}
			if cur.Right != nil {
				q.PushBack(cur.Right)
			}
			// 进入下一层决策树
			dfs(q, track)

			if cur.Left != nil {
				q.Remove(q.Back())
			}
			if cur.Right != nil {
				q.Remove(q.Back())
			}
			q.PushBack(cur)
			track = track[:len(track)-1]
		}

	}
	dfs(q, []int{})
	return ans
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	print(BSTSequences(root))
}
