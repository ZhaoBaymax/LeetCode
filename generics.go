package main

import "fmt"

type Addable interface {
	~int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
}

func add4int(a, b int) int {
	return a + b
}

func add4float32(a, b float32) float32 {
	return a + b
}

func add4string(a, b string) string {
	return a + b
}

func add[T Addable](a, b T) T {
	return a + b
}

func main() {
	a, b := 3, 4
	fmt.Println(add4int(a, b))
	var c, d float32 = 3., 4.
	fmt.Println(add4float32(c, d))
	e, f := "hello", "golang"
	fmt.Println(add4string(e, f))
	fmt.Println()

	fmt.Println(add(a, b))
	fmt.Println(add(c, d))
	fmt.Println(add(e, f))
}
