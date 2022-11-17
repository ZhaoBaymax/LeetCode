package main

import "fmt"

type CustomStruct struct {
	A int
}

type Options func(*CustomStruct)

func WithAddValue(value int) Options {
	return func(c *CustomStruct) {
		c.A += value
	}
}
func WithMinusValue(value int) Options {
	return func(c *CustomStruct) {
		c.A -= value
	}
}
func NewCustomStruct(initValue int, opts ...Options) *CustomStruct {
	res := &CustomStruct{
		A: initValue,
	}

	for _, opt := range opts {
		opt(res)
	}
	return res
}

func main() {

	stru := NewCustomStruct(5, WithAddValue(-1), WithAddValue(2), WithMinusValue(2))
	fmt.Println(stru.A)
}
