package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	length := 0
	h := head
	for head != nil {
		length++
		head = head.Next
	}
	var buildBST func(start, end int) *TreeNode
	buildBST = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) / 2
		left := buildBST(start, mid-1)
		root := &TreeNode{Val: h.Val}
		h = h.Next
		root.Left = left
		root.Right = buildBST(mid+1, end)
		return root
	}
	return buildBST(0, length-1)
}

func main() {
	head := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  9,
							Next: nil,
						},
					},
				},
			},
		},
	}
	sortedListToBST(head)
}
