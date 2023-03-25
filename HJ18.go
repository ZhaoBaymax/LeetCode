package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ipaddr := make([]string, 0)
	for {
		scanner.Scan()
		data := scanner.Text()
		if len(data) == 0 {
			break
		}
		ipaddr = append(ipaddr, data)
	}
	ans := make([]int, 7)
	for i := 0; i < len(ipaddr); i++ {
		flag, flag1, flag2 := check(ipaddr[i])
		if flag == -2 {
			continue
		}
		if flag == -1 || !flag2 {
			ans[5]++
			continue
		} else {
			if flag == 0 {
			}
			ans[flag]++
		}
		if flag1 {
			ans[6]++
		}
	}
	for i := 0; i < len(ans); i++ {
		fmt.Printf("%d ", ans[i])
	}
}

func check(s string) (int, bool, bool) {
	flag := -1
	flag1 := false
	flag2 := true
	str := strings.Split(s, "~")
	ip := net.ParseIP(str[0])
	mask := net.ParseIP(str[1])
	if ip != nil {
		ip = ip.To4()
		if ip[0] >= 240 {
			flag = 4
		} else if ip[0] >= 224 {
			flag = 3
		} else if ip[0] >= 192 {
			flag = 2
		} else if ip[0] >= 128 {
			flag = 1
		} else if ip[0] == 127 || ip[0] == 0 {
			flag = -2
		} else {
			flag = 0
		}
		if ip[0] == 10 || (ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31) || (ip[0] == 192 && ip[1] == 168) {
			flag1 = true
		}
	}
	mask_s := net.IPMask(mask.To4())
	if mask_s == nil {
		flag2 = false
	} else {
		mask_t, _ := strconv.ParseInt(mask_s.String(), 16, 64)
		mask_2 := strconv.FormatInt(mask_t, 2)
		pos_0 := strings.Index(mask_2, "0")
		pos_1 := strings.LastIndex(mask_2, "1")
		if pos_1 == -1 {
			flag2 = false
		} else {
			flag2 = pos_0 >= pos_1
		}
	}
	return flag, flag1, flag2
}
