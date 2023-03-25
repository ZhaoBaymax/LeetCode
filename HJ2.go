package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	var t string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()
	scanner.Scan()
	t = scanner.Text()
	s = strings.ToLower(s)
	fmt.Println(strings.Count(s, strings.ToLower(t)))
}
