package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < T; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(test(a, b, n))
	}
}
func test(a, b, n int) int {
	pre := 2
	cur := a
	for i := 2; i <= n; i++ {
		pre, cur = cur, (cur*a-pre*b)%(1e9+7)
	}
	return (cur + (1e9 + 7)) % (1e9 + 7)
}
