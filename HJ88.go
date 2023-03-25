package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	ss := strings.Split(s, "-")
	a := strings.Split(ss[0], " ")
	b := strings.Split(ss[1], " ")
	test(a, b)
}
func test(a, b []string) {
	paiList := map[string]int{
		"3":     1,
		"4":     2,
		"5":     3,
		"6":     4,
		"7":     5,
		"8":     6,
		"9":     7,
		"10":    8,
		"J":     9,
		"Q":     10,
		"K":     11,
		"A":     12,
		"2":     13,
		"joker": 14,
		"JOKER": 15,
	}
	a_map := map[string]int{}
	b_map := map[string]int{}
	for _, aa := range a {
		a_map[aa]++
	}
	for _, bb := range b {
		b_map[bb]++
	}

	if len(a_map) == len(b_map) && a_map[a[0]] == b_map[b[0]] {
		if paiList[a[0]] > paiList[b[0]] {
			pprint(a)
		} else {
			pprint(b)
		}
	} else {
		if len(a_map) == 2 && a[0] == "joker" {
			pprint(a)
		} else if len(b_map) == 2 && b[0] == "joker" {
			pprint(b)
		} else if len(a) == 4 && len(a_map) == 1 {
			pprint(a)
		} else if len(b) == 4 && len(b_map) == 1 {
			pprint(b)
		} else {
			fmt.Print("ERROR")
		}
	}

}
func pprint(a []string) {
	for _, aa := range a {
		fmt.Print(aa + " ")
	}
}
