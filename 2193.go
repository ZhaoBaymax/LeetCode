package main

import (
	"fmt"
	"strings"
)

func minMovesToMakePalindrome(s string) int {
	ans := 0
	ss := strings.Split(s, "")
	for len(ss) > 1 {
		j := len(ss) - 1
		for k := j; k > 0; k-- {
			if ss[0] == ss[k] {
				for ; k < j; k++ {
					ss[k], ss[k+1] = ss[k+1], ss[k]
					ans++
				}
				ss = ss[1:j]
				break
			}
			if k == 1 {
				ans += len(ss) / 2
				ss = ss[1:]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minMovesToMakePalindrome("scpcyxprxxsjyjrww"))
}
