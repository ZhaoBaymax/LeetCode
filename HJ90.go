package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	s := strings.Split(str, ".")
	if len(s) != 4 {
		fmt.Print("NO")
	} else {
		fmt.Print(test(s))
	}

}

func test(s []string) string {
	for _, v := range s {
		if len(v) > 1 && strings.HasPrefix(v, "0") {
			return "NO"
		}
		ss, err := strconv.Atoi(v)
		if err != nil {
			return "NO"
		}
		if v[0] < '0' || v[0] > '9' {
			return "NO"
		}
		if ss < 0 || ss > 255 {
			return "NO"
		}
	}
	return "YES"
}
