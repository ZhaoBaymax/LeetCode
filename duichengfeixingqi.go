package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	x    int
	y    int
	step int
	k    int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	grid := make([][]string, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		ss := strings.Split(s, "")
		grid = append(grid, ss)
	}
	start := [2]int{}
	end := [2]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == "S" {
				start[0] = i
				start[1] = j
			} else if grid[i][j] == "E" {
				end[0] = i
				end[1] = j
			}
		}
	}
	shortDis(grid, start, end, 5)
}
func shortDis(grid [][]string, pos, end [2]int, k int) {
	n := len(grid)
	m := len(grid[0])
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	e_k := make([][]int, n)
	for i := 0; i < n; i++ {
		e_k[i] = make([]int, m)
		for j := 0; j < m; j++ {
			e_k[i][j] = -1
		}
	}
	start := Node{
		x:    pos[0],
		y:    pos[1],
		step: 0,
		k:    k,
	}
	q := make([]Node, 0)
	q = append(q, start)
	for len(q) != 0 {
		cur_node := q[0]
		q = q[1:]
		if cur_node.x == end[0] && cur_node.y == end[1] {
			fmt.Println(cur_node.step)
			return
		}
		for i := 0; i < 5; i++ {
			next_node := Node{
				x:    0,
				y:    0,
				step: 0,
				k:    0,
			}
			if i < 4 {
				next_node = Node{
					x:    cur_node.x + dirs[i][0],
					y:    cur_node.y + dirs[i][1],
					step: cur_node.step,
					k:    cur_node.k,
				}

			} else {
				next_node = Node{
					x:    n - cur_node.x - 1,
					y:    m - cur_node.y - 1,
					step: cur_node.step,
					k:    cur_node.k - 1,
				}
			}
			if !check(grid, next_node, n, m) {
				continue
			}
			if next_node.k > e_k[next_node.x][next_node.y] {
				next_node.step++
				e_k[next_node.x][next_node.y] = next_node.k
				q = append(q, next_node)
			}
		}
	}
	fmt.Println(-1)
}

func check(grid [][]string, node Node, n, m int) bool {
	if !(0 <= node.x && node.x < n && 0 <= node.y && node.y < m) {
		return false
	}
	if node.k < 0 {
		return false
	}
	if grid[node.x][node.y] == "#" {
		return false
	}
	return true
}
