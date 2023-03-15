package main

import "container/list"

type MinStack struct {
	minElement []int
	ss         *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	l := list.New()
	stack := MinStack{
		minElement: []int{},
		ss:         l,
	}
	return stack
}

func (this *MinStack) Push(x int) {
	if this.IsEmpty() || x < this.minElement[len(this.minElement)-1] {
		this.minElement = append(this.minElement, x)
	} else {
		this.minElement = append(this.minElement, this.minElement[len(this.minElement)-1])
	}
	this.ss.PushFront(x)
}

func (this *MinStack) Pop() {
	if !this.IsEmpty() {
		this.minElement = this.minElement[:len(this.minElement)-2]
		this.ss.Remove(this.ss.Front())
	}
}

func (this *MinStack) Top() int {
	if !this.IsEmpty() {
		return this.ss.Front().Value.(int)
	}
	return -1
}

func (this *MinStack) GetMin() int {
	return this.minElement[len(this.minElement)-1]
}

func (this *MinStack) IsEmpty() bool {
	return this.ss.Len() == 0
}
