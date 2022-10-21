package main

import (
	"fmt"
)

//
//type Point struct {
//	x int
//	y int
//}

// BFS 超时了！！！
//func minPathSum(grid [][]int) int {
//	m := len(grid)
//	n := len(grid[0])
//	track := make([][]int, m)
//	for i := range track {
//		track[i] = make([]int, n)
//	}
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			track[i][j] = math.MaxInt
//		}
//	}
//	track[0][0] = grid[0][0]
//	dirs := []Point{{0, 1}, {1, 0}}
//	q := []Point{}
//	q = append(q, Point{0, 0})
//	for len(q) > 0 {
//		cur_pos := q[0]
//		q = q[1:]
//		for i := 0; i < 2; i++ {
//			x := cur_pos.x
//			y := cur_pos.y
//			dist := track[x][y]
//			next := Point{x + dirs[i].x, y + dirs[i].y}
//			if !(0 <= next.x && next.x < m && next.y < n && 0 <= next.y) {
//				continue
//			}
//			dist = dist + grid[next.x][next.y]
//			if track[next.x][next.y] >= dist {
//				track[next.x][next.y] = dist
//				q = append(q, next)
//			}
//		}
//	}
//	return track[m-1][n-1]
//}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i][j-1]+grid[i][j], dp[i-1][j]+grid[i][j])
			}
		}
	}
	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}
func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	//grid := [][]int{{0, 0}, {0, 0}}
	fmt.Println(minPathSum(grid))
}
