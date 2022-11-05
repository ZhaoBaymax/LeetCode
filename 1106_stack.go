package main

import (
	"container/list"
	"fmt"
)

func parseBoolExpr(expression string) bool {
	stack := list.New()
	for _, c := range expression {
		if c == ',' {
			continue
		}
		if c != ')' {
			stack.PushBack(c)
			continue
		}
		t_f_reflect := make(map[int32]int, 2)
		for top := stack.Back(); top.Value.(int32) != '('; {
			t_f_reflect[top.Value.(int32)]++
			top = top.Prev()
			stack.Remove(stack.Back())
			if top.Value.(int32) == '(' {
				op := top.Prev().Value.(int32)
				stack.Remove(stack.Back())
				stack.Remove(stack.Back())
				if op == '!' {
					if _, ok := t_f_reflect['f']; ok {
						stack.PushBack('t')
					} else {
						stack.PushBack('f')
					}
				} else if op == '|' {
					if _, ok := t_f_reflect['t']; ok {
						stack.PushBack('t')
					} else {
						stack.PushBack('f')
					}
				} else {
					if _, ok := t_f_reflect['f']; ok {
						stack.PushBack('f')
					} else {
						stack.PushBack('t')
					}
				}
				break
			}
		}
	}
	return stack.Back().Value.(int32) == 't'
}
func main() {
	expression := "|(&(t,f,t),!(t))"
	fmt.Println(parseBoolExpr(expression))
}
