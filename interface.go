package main

import "fmt"

type tt interface {
	say()
}

func main() {
	var i tt
	str := ""
	if i == nil {
		fmt.Println("nil", str)
		return
	}
	fmt.Println("not nil")
}
