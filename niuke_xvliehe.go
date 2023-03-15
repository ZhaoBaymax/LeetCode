package main

import "fmt"

func main() {
	a := 0
	b := 0

	fmt.Scan(&a, &b)

	start, ansLen := seqSum(a, b)
	if ansLen == a {
		fmt.Println("No")
	} else {
		for i := start; i < ansLen+start; i++ {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}
}

func seqSum(n, k int) (start, ansLen int) {
	left, right := 0, 0
	ansLen = n
	sum := 0
	for right < n/2 {
		sum += right
		right++
		if right-left > 100 {
			sum = sum - left
			left++
		}
		for sum > n {
			sum = sum - left
			left++
		}
		if sum == n && right-left >= k {
			if right-left < ansLen {
				ansLen = right - left
				start = left
				if left == 0 && right-left-1 >= k {
					ansLen--
					start++
				}
			}
			sum = sum - left
			left++
		}
	}
	return
}
