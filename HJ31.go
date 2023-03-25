package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	var splitFunc func(r rune) bool
	splitFunc = func(r rune) bool {
		if ('a' <= r && r <= 'z') || 'A' <= r && r <= 'Z' {
			return false
		}
		return true
	}
	s := strings.FieldsFunc(str, splitFunc)
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Print(s[i])
		fmt.Print(" ")
	}
}
