package main

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}
func (stack Stack) Cap() int {
	return cap(stack)
}
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

func reverseParentheses(s string) string {
	var stack Stack
	var cur_s []byte
	for i := range s {
		if s[i] == '(' {
			stack.Push(cur_s)
			cur_s = []byte{}
		} else if s[i] == ')' {
			for j, k := 0, len(cur_s); j < k/2; j++ {
				cur_s[j], cur_s[k-j-1] = cur_s[k-j-1], cur_s[j]
			}
			t, _ := stack.Pop()
			cur_s = append(t.([]byte), cur_s...)
		} else {
			cur_s = append(cur_s, s[i])
		}
	}
	return string(cur_s)
}
func main() {
	fmt.Println(reverseParentheses("(u(love)i)"))
}
