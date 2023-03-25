package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	strs := make([]string, 0)
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		strs = append(strs, scanner.Text())
	}
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	num, ss := getBrother(strs, s, k)
	fmt.Println(num)
	fmt.Println(ss)
}
func getBrother(strs []string, s string, k int) (int, string) {
	val := map[int32]int{}
	need := map[int32]int{}
	for _, v := range s {
		val[v]++
	}
	count := 0
	valid := 0
	ans := make([]string, 0)
	for i := 0; i < len(strs); i++ {
		if strs[i] == s || len(strs[i]) != len(s) {
			continue
		}
		for _, v := range strs[i] {
			if _, ok := val[v]; ok {
				need[v]++
				if need[v] == val[v] {
					valid++
				}
			} else {
				break
			}
		}
		if valid == len(val) {
			count++
			ans = append(ans, strs[i])
		}
		valid = 0
		need = map[int32]int{}
	}
	sort.SliceStable(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	if count == 0 || k > len(ans) {
		return count, ""
	}
	return count, ans[k-1]
}
