package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := n / 2; i > 0; i-- {
		if check(i) && check(n-i) {
			fmt.Println(i)
			fmt.Println(n - i)
			break
		}
	}
}

func check(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%6 != 1 && n%6 != 5 {
		return false
	}
	sq := math.Sqrt(float64(n))
	for i := 6; i <= int(sq)+1; i = i + 6 {
		if n%(i-1) == 0 || n%(i+1) == 0 {
			return false
		}
	}
	return true
}
