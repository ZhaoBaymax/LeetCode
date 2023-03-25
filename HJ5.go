package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)
	num, _ := strconv.ParseInt(str, 0, 0)
	fmt.Println(num)
}
