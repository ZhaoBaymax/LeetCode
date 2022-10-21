package main

import "fmt"

func main() {
	s := make([]int, 5)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}