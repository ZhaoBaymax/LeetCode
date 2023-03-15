package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	hair := &ListNode{
		Val:  0,
		Next: head,
	}

	p := head
	q := head.Next
	for q != nil {
		if p.Val >= q.Val {
			p = p.Next
		} else {
			pre := hair
			for pre.Next.Val <= q.Val {
				pre = pre.Next
			}
			p.Next = q.Next
			q.Next = pre.Next
			pre.Next = q
		}
		q = p.Next
	}
	return hair.Next
}

func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	sortList(head)
}
