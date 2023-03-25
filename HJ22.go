package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	nums := make([]int, 0)
	for {
		scanner.Scan()
		data := scanner.Text()
		if data == "0" {
			break
		}
		num, _ := strconv.Atoi(data)
		nums = append(nums, num)
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(drink(nums[i]))
	}
}
func drink(n int) string {
	num := big.NewInt(int64(n))
	sum := big.NewInt(0)
	res := num
	for res.Int64() >= 3 {
		div, mod := res.DivMod(res, big.NewInt(3), big.NewInt(0))
		sum = sum.Add(sum, div)
		res = res.Add(div, mod)
	}
	if res.Int64() == 2 {
		sum = sum.Add(sum, big.NewInt(1))
	}
	return sum.String()
}
