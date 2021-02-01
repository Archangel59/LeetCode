package main

import "fmt"

var vis = [60][60]bool{}
var dir = [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func dfs(grid [][]int, n, m, x, y, ans int) {
	if x >= 0 && x < n && y >= 0 && y < m {
		if !vis[x][y] && grid[x][y] <= ans {
			vis[x][y] = true
			for i := 0; i < 4; i++ {
				dfs(grid, n, m, x+dir[i][0], y+dir[i][1], ans)
			}
		}
	}
}

func check(grid [][]int, n, m, ans int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			vis[i][j] = false
		}
	}
	dfs(grid, n, m, 0, 0, ans)
	return vis[n-1][m-1]
}

func swimInWater(grid [][]int) int {
	l, r, n, m := 1000, 0, len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := grid[i][j]
			if l > x {
				l = x
			}
			if r < x {
				r = x
			}
		}
	}

	for {
		if l >= r {
			break
		}
		mid := (l + r) / 2
		if check(grid, n, m, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func main() {
	fmt.Println(swimInWater([][]int{{3, 2}, {0, 1}}))
}
