package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)
	num, _ := strconv.ParseFloat(str, 64)
	num = num + 0.5
	fmt.Println(int(num))
}
