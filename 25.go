package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	p := newHead // work pointer
	for head != nil {
		tail := p
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return newHead.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		//newHead.Next = p
		p.Next = head
		tail.Next = nex
		p = tail
		head = tail.Next
	}
	return newHead.Next
}

func myReverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	nex := tail.Next
	p := head
	for nex != tail {
		temp := p.Next
		p.Next = nex
		nex = p
		p = temp
	}
	return tail, head
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
	reverseKGroup(head, 3)
}
