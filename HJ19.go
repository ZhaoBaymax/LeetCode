package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	path_info := make([]string, 0)
	for {
		scanner.Scan()
		data := scanner.Text()
		if len(data) == 0 {
			break
		}
		path_info = append(path_info, data)
	}
	ll := list.New()
	path_map := map[string]int{}
	for i := 0; i < len(path_info); i++ {
		path := strings.Split(path_info[i], " ")
		path_last := strings.Split(path[0], "\\")
		path_l := path_last[len(path_last)-1]
		if len(path_l) > 16 {
			path_l = path_l[len(path_l)-16:]
		}
		path_l = path_l + " " + path[1]
		if _, ok := path_map[path_l]; ok {
			path_map[path_l]++
		} else {
			path_map[path_l]++
			ll.PushBack(path_l)
		}
	}
	if ll.Len() <= 8 {
		print_ans(ll.Front(), path_map)
	}
	for l, p, count := ll.Front(), ll.Front(), 0; l != nil; l, count = l.Next(), count+1 {
		if count < 8 {
			continue
		} else {
			p = p.Next()
		}
		if l.Next() == nil {
			print_ans(p, path_map)
		}
	}

}

func print_ans(p *list.Element, path_map map[string]int) {
	for ; p != nil; p = p.Next() {
		cnt := path_map[p.Value.(string)]
		pp := strings.Split(p.Value.(string), " ")
		fmt.Printf("%s %s %d\n", pp[0], pp[1], cnt)
	}
}
