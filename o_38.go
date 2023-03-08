package main

import (
	"fmt"
	"sort"
)

func permutation(s string) []string {
	var ans []string
	byteS := []byte(s)
	sort.Slice(byteS, func(i, j int) bool {
		return byteS[i] < byteS[j]
	})
	used := make([]bool, len(s))
	track := make([]byte, 0, len(s))
	var trackback func(byteS []byte)
	trackback = func(byteS []byte) {
		if len(track) == len(byteS) {
			ans = append(ans, string(track))
			return
		}
		for k, ss := range byteS {
			if !used[k] {
				if k >= 1 && ss == byteS[k-1] && !used[k-1] {
					continue
				}
				track = append(track, ss)
				used[k] = true
				trackback(byteS)
				used[k] = false
				track = track[:len(track)-1]
			}
		}
	}
	trackback(byteS)
	return ans
}

func isContain(ss byte, track []byte) bool {
	for _, v := range track {
		if ss == v {
			return true
		}
	}
	return false
}
func isContainString(a string, b []string) bool {
	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}

func main() {
	s := "aab"
	fmt.Println(permutation(s))
}
