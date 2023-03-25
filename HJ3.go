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
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, 501)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		nums[num]++
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			fmt.Println(i)
		}
	}
}
