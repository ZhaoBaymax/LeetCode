package main

import "strconv"

func maxValue(n string, x int) string {
	l := len(n)
	i := 0
	if n[0] == '-' {
		for i = 1; i < l; i++ {
			if int(n[i]-'0') >= x {
				break
			}
		}
	} else {
		for i = 0; i < l; i++ {
			if int(n[i]-'0') <= x {
				break
			}
		}
	}
	return n[:i] + strconv.Itoa(x) + n[i:]
}

func main() {
	n := "1199"
	x := 9
	print(maxValue(n, x))
}
