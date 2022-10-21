package main

type S int

func main() {
	a := S(0)
	b := make([]*S, 2)
	b[0] = &a
	c := new(S)
	b[1] = c
}
