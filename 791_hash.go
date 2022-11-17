package main

import "sort"

func customSortString(order string, s string) string {
	m := map[byte]int{}
	for i := 0; i < len(order); i++ {
		m[order[i]] = i + 1
	}
	ans := []byte(s)
	sort.Slice(ans, func(i, j int) bool {
		return m[ans[i]] < m[ans[j]]
	})
	return string(ans)
}
