package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	var y int
	var a int
	var b int
	fmt.Scan(&x, &y, &a, &b)
	g := gbc(a, b)
	a = a / g
	b = b / g
	u := min(x/a, y/b)
	fmt.Println(strconv.Itoa(u*a) + " " + strconv.Itoa(u*b))
}

func gbc(x, y int) int {
	if x < y {
		x, y = y, x
	}
	if y == 0 {
		return x
	}
	return gbc(y, x%y)
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
