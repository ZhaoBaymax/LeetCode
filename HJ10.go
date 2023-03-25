package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(test(s))
}

func test(s string) int {
	ss_map := map[int32]int{}
	for _, s_byte := range s {
		if s_byte >= 0 && s_byte <= 127 {
			ss_map[s_byte]++
		}
	}
	return len(ss_map)
}
