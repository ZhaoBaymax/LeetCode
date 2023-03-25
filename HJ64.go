package main

import "fmt"

func main() {
	var n int
	var command string
	fmt.Scan(&n, &command)
	mp3(n, command)
}

func mp3(n int, command string) {
	c := map[uint8]int{
		'U': -1,
		'D': 1,
	}
	if n <= 3 {
		start := 1
		for i := 0; i < len(command); i++ {
			start = start + c[command[i]]
			if start == 0 {
				start = n
			} else if start > n {
				start = 1
			}
		}
		for i := 1; i <= n; i++ {
			fmt.Print(i)
			fmt.Print(" ")
		}
		fmt.Println()
		fmt.Print(start)
	} else {
		idx := 1
		start := 1
		end := 4
		for i := 0; i < len(command); i++ {
			idx = idx + c[command[i]]
			if idx < start {
				if start == 1 {
					idx = n
					start = n - 3
					end = n
				} else {
					start--
					end--
				}
			} else if idx > end {
				if end == n {
					idx = 1
					start = 1
					end = 4
				} else {
					start++
					end++
				}
			}
		}
		for i := start; i <= end; i++ {
			fmt.Print(i)
			fmt.Print(" ")
		}
		fmt.Println()
		fmt.Print(idx)
	}
}
