package main

import (
	"fmt"
	"math"
)

func maxTurbulenceSize(arr []int) int {
	len_arr := len(arr)
	if len_arr == 1 {
		return len_arr
	}
	cur_max_u := make([]int, len_arr)
	cur_max_d := make([]int, len_arr)
	max := 0
	for i := 1; i < len_arr; i++ {
		if arr[i] > arr[i-1] {
			cur_max_u[i] = cur_max_d[i-1] + 1
		} else if arr[i] < arr[i-1] {
			cur_max_d[i] = cur_max_u[i-1] + 1
		}
		max = int(math.Max(float64(max), math.Max(float64(cur_max_u[i]), float64(cur_max_d[i]))))
	}
	return max + 1
}

func main() {
	arr := []int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24}
	fmt.Println(maxTurbulenceSize(arr))
}
