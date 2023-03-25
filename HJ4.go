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
	str := scanner.Text()
	build := strings.Builder{}
	ans := make([]strings.Builder, 0)
	i := 0
	for ; i < len(str); i++ {
		build.WriteByte(str[i])
		if i%8 == 7 {
			ans = append(ans, build)
			build.Reset()
		}
	}
	if len(str)%8 != 0 {
		for i = 0; i < 8-(len(str)%8); i++ {
			build.WriteByte('0')
		}
	}
	ans = append(ans, build)
	for i = 0; i < len(ans); i++ {
		fmt.Println(ans[i].String())
	}

}
