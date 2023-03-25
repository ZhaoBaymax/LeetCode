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
	T, _ := strconv.Atoi(scanner.Text())
	t := make([][]int, T+1)
	for i := 1; i <= T; i++ {
		t[i] = make([]int, 2)
		scanner.Scan()
		t[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		t[i][1], _ = strconv.Atoi(scanner.Text())
	}
	ans := 0
	hight := make([]int, T+1)
	for i := T; i >= 1; i-- {
		if t[i][0] == -1 && t[i][1] == -1 {
			hight[i] = 1
			ans++
			continue
		}
		if t[i][1] != -1 && hight[t[i][0]] == hight[t[i][1]] {
			hight[i] = hight[t[i][0]] + 1
			ans++
		}
	}
	fmt.Println(ans)
}

//
//func hi(t [][]int, start int) int {
//	if t[start][0] == -1 && t[start][1] == -1 {
//		return 1
//	}
//	return max(hi(t, t[start][0]), hi(t, t[start][1])) + 1
//}
//
//func test(t [][]int, start int) bool {
//	if t[start][0] == -1 && t[start][1] == -1 {
//		return true
//	}
//	if test(t, t[start][0]) && test(t, t[start][1]) && hi(t, t[start][0]) == hi(t, t[start][1]) {
//		return true
//	}
//	return false
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
