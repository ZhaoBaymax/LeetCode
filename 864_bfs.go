package main

import (
	"container/list"
	"fmt"
	"unicode"
)

type Point struct {
	x    int
	y    int
	mask int
	l    int
	step int
}

func shortestPathAllKeys(grid []string) int {
	m := len(grid)
	n := len(grid[0])
	key := 0
	keyToIdx := map[rune]int{}
	track := make([][]Point, m)
	for i := 0; i < m; i++ {
		track[i] = make([]Point, n)
	}
	for i, row := range track {
		for j, _ := range row {
			track[i][j] = Point{
				x:    i,
				y:    j,
				mask: -1,
				l:    0,
				step: 0,
			}
		}
	}
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	q := list.New()
	for i, row := range grid {
		for j, c := range row {
			if c == '@' {
				start := Point{
					x:    i,
					y:    j,
					mask: 0,
					l:    0,
					step: 0,
				}
				q.PushBack(start)
				track[i][j] = start
			} else if unicode.IsLower(c) {
				keyToIdx[c] = key
				key++
			}
		}
	}

	for q.Len() > 0 {
		cur := q.Remove(q.Front()).(Point)
		for i := 0; i < 4; i++ {
			next_point := Point{
				x:    cur.x + dirs[i][0],
				y:    cur.y + dirs[i][1],
				mask: cur.mask,
				l:    cur.l,
				step: cur.step,
			}
			if next_point.x < 0 || next_point.x >= m || next_point.y < 0 || next_point.y >= n || grid[next_point.x][next_point.y] == '#' {
				continue
			}

			if unicode.IsLower(rune(grid[next_point.x][next_point.y])) {
				mask := next_point.mask | 1<<keyToIdx[rune(grid[next_point.x][next_point.y])]
				if mask != next_point.mask { // 新的钥匙
					next_point.l++
					next_point.mask = mask
				} else {
					continue
				}
				if next_point.l == key {
					return next_point.step + 1
				}
			}
			if unicode.IsUpper(rune(grid[next_point.x][next_point.y])) {
				mask := next_point.mask | 1<<keyToIdx[unicode.ToLower(rune(grid[next_point.x][next_point.y]))]
				if mask != next_point.mask { // 没有钥匙
					continue
				}
			}
			next_point.step++
			track[next_point.x][next_point.y] = next_point
			q.PushBack(next_point)
		}
	}
	return -1
}

//func equal(x, y map[byte]int) bool {
//	if x == nil || y == nil {
//		return false
//	}
//	//if len(x)  == 0 && len(y) == 0{
//	//	return false
//	//}
//	if len(x) != len(y) {
//		return false
//	}
//	for k, xv := range x {
//		if yv, ok := y[k]; !ok || yv != xv {
//			return false
//		}
//	}
//	return true
//}

func main() {
	grid := []string{"Dd#b@", ".fE.e", "##.B.", "#.cA.", "aF.#C"}
	fmt.Println(shortestPathAllKeys(grid))
}
