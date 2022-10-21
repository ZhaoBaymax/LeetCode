package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	l := 0
	top := 0
	r := n - 1
	b := n - 1
	num := 1
	for num <= n*n {
		for i := l; i <= r && num <= n*n; i++ {
			res[top][i] = num
			num++
		}
		top++
		for i := top; i <= b && num <= n*n; i++ {
			res[i][r] = num
			num++
		}
		r--
		for i := r; i >= l && num <= n*n; i-- {
			res[b][i] = num
			num++
		}
		b--
		for i := b; i >= top && num <= n*n; i-- {
			res[i][l] = num
			num++
		}
		l++
	}
	return res

}

func main() {
	fmt.Println(generateMatrix(3))
}
