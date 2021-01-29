package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	return max(a, -a)
}

var vis = [105][105]bool{}
var dir = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func dfs(heights [][]int, x, y, num int) {
	n, m := len(heights), len(heights[0])
	if vis[n-1][m-1] {
		return
	}
	for i := 0; i < 4; i++ {
		nx, ny := x+dir[i][0], y+dir[i][1]
		if nx >= 0 && nx < n && ny >= 0 && ny < m && !vis[nx][ny] {
			if abs(heights[nx][ny]-heights[x][y]) <= num {
				vis[nx][ny] = true
				dfs(heights, nx, ny, num)
			}
		}
	}
}

func minimumEffortPath(heights [][]int) int {
	l, r, n, m := 0, 0, len(heights), len(heights[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			r = max(r, heights[i][j])
		}
	}
	for {
		if l >= r {
			return r
		}
		mid := (l + r) >> 1
		check := func() bool {
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					vis[i][j] = false
				}
			}
			vis[0][0] = true
			dfs(heights, 0, 0, mid)
			return vis[n-1][m-1]
		}
		if check() {
			r = mid
		} else {
			l = mid + 1
		}
	}
}

func main() {
	fmt.Println(minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}))
}
