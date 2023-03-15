package main

import (
	"fmt"
)

func deleteS(s string, n, k int) string {
	byte_s := []byte(s)
	stack := make([]byte, 0)
	count := 0
	for i := 0; i < n; i++ {
		for len(stack) != 0 && stack[len(stack)-1] > byte_s[i] && count < k {
			stack = stack[:len(stack)-1]
			count++
		}
		stack = append(stack, byte_s[i])
	}
	for ; count < k; count++ {
		stack = stack[:len(stack)-1]
	}
	return string(stack)
}

func main() {
	var a int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		var n, k int
		fmt.Scan(&n, &k)
		var s string
		fmt.Scan(&s)
		fmt.Println(deleteS(s, n, k))
	}
}
