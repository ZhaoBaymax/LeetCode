package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, T)
	for i := 0; i < T; i++ {
		scanner.Scan()
		nums[i], _ = strconv.Atoi(scanner.Text())
	}
	odd := make([]int, 0)
	even := make([]int, 0)
	for i := 0; i < T; i++ {
		if nums[i]%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	g := make([][]int, len(odd))
	for i := 0; i < len(odd); i++ {
		g[i] = make([]int, len(even))
		for j := 0; j < len(even); j++ {
			if isPrime(odd[i] + even[j]) {
				g[i][j] = odd[i] + even[j]
			}
		}
	}
	visited := make([]bool, len(even))
	link := make([]int, len(even))
	for i := 0; i < len(link); i++ {
		link[i] = -1
	}
	var dfs func(i int) bool
	dfs = func(i int) bool {
		for j := 0; j < len(even); j++ {
			if g[i][j] != 0 && !visited[j] {
				visited[j] = true
				if link[j] == -1 || dfs(link[j]) {
					link[j] = i
					return true
				}
			}
		}
		return false
	}
	ans := 0
	for i := 0; i < len(odd); i++ {
		if dfs(i) {
			ans++
		}
		visited = make([]bool, len(even))
	}
	fmt.Println(ans)
}

func isPrime(num int) bool {
	if num <= 3 {
		return num > 1
	}
	if num%6 != 5 && num%6 != 1 {
		return false
	}
	sq := int(math.Sqrt(float64(num)))
	for i := 5; i <= sq; i = i + 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
