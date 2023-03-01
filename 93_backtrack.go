package main

import "strconv"

var (
	ans      []string
	segments []int
)

func restoreIpAddresses(s string) []string {
	segments = make([]int, 4)
	ans = []string{}
	dfs(s, 0, 0)
	return ans
}

func dfs(s string, seg, start int) {
	if seg == 4 {
		if start == len(s) {
			ipaddr := ""
			for i := 0; i < 4; i++ {
				ipaddr += strconv.Itoa(segments[i])
				if i != 3 {
					ipaddr += "."
				}
			}
			ans = append(ans, ipaddr)
		}
		return
	}
	if start == len(s) {
		return
	}
	if s[start] == '0' {
		segments[seg] = 0
		dfs(s, seg+1, start+1)
	}
	addr := 0
	for segEnd := start; segEnd < len(s); segEnd++ {
		addr = addr*10 + int(s[segEnd]-'0')
		if addr > 0 && addr <= 0xFF {
			segments[seg] = addr
			dfs(s, seg+1, segEnd+1)
		} else {
			break
		}
	}
}
