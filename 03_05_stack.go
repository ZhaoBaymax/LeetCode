package main

import (
	"container/list"
	"fmt"
)

type SortedStack struct {
	stack *list.List
}

func Constructor() SortedStack {
	stack := list.New()
	sortedStack := &SortedStack{stack: stack}
	return *sortedStack
}

func (this *SortedStack) Push(val int) {
	if this.IsEmpty() {
		this.stack.PushBack(val)
	} else {
		i := this.stack.Front()
		for ; i != this.stack.Back(); i = i.Next() {
			if val > i.Value.(int) {
				this.stack.InsertBefore(val, i)
				return
			}
		}
		if i == this.stack.Back() {
			if val < this.stack.Back().Value.(int) {
				this.stack.InsertAfter(val, i)
			} else {
				this.stack.InsertBefore(val, i)
			}
		}
	}
}

func (this *SortedStack) Pop() {
	if !this.IsEmpty() {
		this.stack.Remove(this.stack.Back())
	}
}

func (this *SortedStack) Peek() int {
	if !this.IsEmpty() {
		return this.stack.Back().Value.(int)
	}
	return -1
}

func (this *SortedStack) IsEmpty() bool {
	return this.stack.Len() == 0
}
func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Peek())
	obj.Pop()
	fmt.Println(obj.Peek())
}
