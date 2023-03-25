package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ip := scanner.Text()
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	fmt.Println(parseIPToInt(ip))
	fmt.Println(parseIPToString(s))

}

func parseIPToInt(s string) int64 {
	ip := net.ParseIP(s)
	ip = ip.To4()
	first := int64(ip[0]) << 24
	second := int64(ip[1]) << 16
	third := int64(ip[2]) << 8
	four := int64(ip[3])
	sum := first + second + third + four
	return sum
}

func parseIPToString(num int) string {
	ipv4 := make(net.IP, 4)
	for i := 0; i < 4; i++ {
		ipv4[3-i] = byte(num & 0xFF)
		num = num >> 8
	}
	return ipv4.String()
}
