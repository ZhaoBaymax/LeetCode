package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	l := list.New()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	head, _ := strconv.Atoi(scanner.Text())
	l.PushBack(head)
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		n1, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		n2, _ := strconv.Atoi(scanner.Text())
		for node := l.Front(); node != nil; node = node.Next() {
			if node.Value.(int) == n2 {
				l.InsertAfter(n1, node)
				break
			}
		}
	}
	scanner.Scan()
	del, _ := strconv.Atoi(scanner.Text())
	for node := l.Front(); node != nil; node = node.Next() {
		if node.Value.(int) == del {
			l.Remove(node)
		}
	}
	for node := l.Front(); node != nil; node = node.Next() {
		fmt.Print(node.Value.(int))
		fmt.Print(" ")
	}
}
