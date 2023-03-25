package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	ans := make([]uint8, len(s))
	word := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if ('a' <= s[i] && s[i] <= 'z') || ('A' <= s[i] && s[i] <= 'Z') {
			word = append(word, s[i])
		} else {
			ans[i] = s[i]
		}
	}

	sort.SliceStable(word, func(i, j int) bool {
		//if strings.ToLower(string(word[i])) == strings.ToLower(string(word[j])) {
		//	return i < j
		//}
		return strings.ToLower(string(word[i])) < strings.ToLower(string(word[j]))
	})
	for i, j := 0, 0; i < len(ans); i++ {
		if ans[i] == 0 {
			ans[i] = word[j]
			j++
		}
	}
	fmt.Println(string(ans))
}
