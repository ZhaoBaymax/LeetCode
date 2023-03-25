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
	str := strings.Split(scanner.Text(), " ")
	num := len(str[len(str)-1])
	fmt.Println(num)
}
