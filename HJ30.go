package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	a := scanner.Text()
	scanner.Scan()
	b := scanner.Text()
	rest(a, b)
}

func rest(a, b string) {
	c := strings.Split(a+b, "")
	ji := make([]string, 0)
	ou := make([]string, 0)
	for i, cc := range c {
		if i%2 == 0 {
			ou = append(ou, cc)
		} else {
			ji = append(ji, cc)
		}
	}
	sort.Slice(ji, func(i, j int) bool {
		return ji[i] < ji[j]
	})
	sort.Slice(ou, func(i, j int) bool {
		return ou[i] < ou[j]
	})
	s := make([]string, 0)
	for i := 0; i < len(ji); i++ {
		s = append(s, ou[i], ji[i])
		if i == len(ji)-1 && i+1 < len(ou) {
			s = append(s, ou[i+1])
		}
	}
	//ans = make([]string, len(c))
	for _, v := range s {
		if strings.ToLower(v) > "f" {
			fmt.Print(v)
			continue
		}
		ten, _ := strconv.ParseInt(strings.ToLower(v), 16, 0)
		bi := strconv.FormatInt(ten, 2)
		re_bi, _ := strconv.ParseInt(reverse(bi), 2, 64)
		re_ten := strings.ToUpper(strconv.FormatInt(re_bi, 16))
		fmt.Print(re_ten)
	}
}

func reverse(s string) string {
	r := []byte(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	for i := len(r); i < 4; i++ {
		r = append(r, '0')
	}
	return string(r)
}
