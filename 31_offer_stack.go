package main

import "container/list"

func validateStackSequences(pushed []int, popped []int) bool {
	q := list.New()
	for i := 0; i < len(pushed); i++ {
		q.PushBack(pushed[i])
		for len(popped) > 0 && q.Len() > 0 {
			if q.Back().Value == popped[0] {
				q.Remove(q.Back())
				popped = popped[1:]
			} else {
				break
			}
		}
	}
	return len(popped) == 0
}

func main() {
	pushed := []int{0, 1}
	popped := []int{0, 1}
	print(validateStackSequences(pushed, popped))
}
