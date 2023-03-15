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
	g := make([][]int, 3)
	for i := 0; i < 3; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			g[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}
	fmt.Print(dji(g))
	//var a int
	//fmt.Scan(&a)
	//g := make([][]int, 3)
	//for i := 0; i < 3; i++ {
	//	for j := 0; j < a; j++ {
	//		var ele int
	//		fmt.Scan(&ele)
	//		g[i] = append(g[i], ele)
	//	}
	//}
	//fmt.Print(dji(g))
}

func dji(g [][]int) int {
	row := len(g)
	col := len(g[0])
	last := make([]int, row)
	for i := 1; i < col; i++ {
		cur_last := make([]int, row)
		for j := 0; j < row; j++ {
			for k := 0; k < row; k++ {
				if k == 0 {
					cur_last[j] = abs(g[j][i], g[k][i-1]) + last[k]
				} else {
					cur_last[j] = min(cur_last[j], abs(g[j][i], g[k][i-1])+last[k])
				}
			}
		}
		last = cur_last
	}
	min_res := last[0]
	for i := 1; i < row; i++ {
		min_res = min(min_res, last[i])
	}
	return min_res
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
