package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	q := head
	i := 0
	for ; p.Next != nil; i++ {
		if i < n {
			p = p.Next
		} else {
			p = p.Next
			q = q.Next
		}
	}
	if i == n-1 {
		return head.Next
	}
	q.Next = q.Next.Next
	return head
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	removeNthFromEnd(head, 1)
}
