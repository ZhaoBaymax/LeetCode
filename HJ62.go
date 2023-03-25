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
	for {
		scanner.Scan()
		data := scanner.Text()
		if len(data) == 0 {
			break
		}
		num, _ := strconv.Atoi(data)
		bi := strconv.FormatInt(int64(num), 2)
		count := strings.Count(bi, "1")
		fmt.Println(count)
	}
}
