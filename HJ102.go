package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Scan(&s)
	countWord(s)
}
func countWord(s string) {
	ha := map[int32]int{}
	for _, ss := range s {
		ha[ss]++
	}
	kk := make([]int32, len(ha))

	i := 0
	for k, _ := range ha {
		kk[i] = k
		i++
	}
	sort.Slice(kk, func(i, j int) bool {
		if ha[kk[i]] == ha[kk[j]] {
			return kk[i] < kk[j]
		}
		return ha[kk[i]] > ha[kk[j]]
	})
	for i = 0; i < len(kk); i++ {
		fmt.Print(string(kk[i]))
	}
}
