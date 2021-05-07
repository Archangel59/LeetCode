package main

import "fmt"

var dir = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func dfs(grid [][]byte, x, y, n, m int) {
	grid[x][y] = '0'
	for i := 0; i < 4; i++ {
		nx, ny := x+dir[i][0], y+dir[i][1]
		if nx >= 0 && nx < n && ny >= 0 && ny < m && grid[nx][ny] == '1' {
			dfs(grid, nx, ny, n, m)
		}
	}
}

func numIslands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(grid, i, j, n, m)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numIslands([][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))
}
