package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	sum := 23.0 * x / 8.0
	fmt.Printf("%0.6f\n", sum)
	a := x / 32.0
	fmt.Printf("%0.6f\n", a)
}
