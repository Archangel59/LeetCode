package main

import "fmt"

var dir = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
var a = [100][100]int{}

func dfs(i, j, n int) {
	if a[i][j] == 1 {
		return
	}
	a[i][j] = 1
	for _, v := range dir {
		x, y := i+v[0], j+v[1]
		if x >= 0 && x < n && y >= 0 && y < n {
			dfs(i+v[0], j+v[1], n)
		}
	}
}

func regionsBySlashes(grid []string) int {
	n := len(grid)
	for i := 0; i <= n*3; i++ {
		for j := 0; j <= n*3; j++ {
			a[i][j] = 0
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x, y := i*3, j*3
			switch grid[i][j] {
			case '/':
				a[x][y+2], a[x+1][y+1], a[x+2][y] = 1, 1, 1
			case '\\':
				a[x][y], a[x+1][y+1], a[x+2][y+2] = 1, 1, 1
			}
		}
	}

	ans := 0
	for i := 0; i < 3*n; i++ {
		for j := 0; j < 3*n; j++ {
			if a[i][j] == 0 {
				ans++
				dfs(i, j, n*3)
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(regionsBySlashes([]string{"/\\", "\\/"}))
}
