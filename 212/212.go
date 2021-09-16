package main

import (
	"fmt"
	"strconv"
	"strings"
)

var M [][]byte // 地图

var dir = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // 搜索方向 上下左右

var flag bool        // 表示当前是否能走完
var vis [15][15]bool // 记录走过的地方

func dfs(x, y, l int, w string) {
	if l+1 == len(w) {
		flag = true
		return
	}
	for i := 0; i < 4; i++ {
		nx, ny := x+dir[i][0], y+dir[i][1]
		if nx >= 0 && nx < len(M) && ny >= 0 && ny < len(M[0]) {
			if M[nx][ny] == w[l+1] && !vis[nx][ny] {
				vis[nx][ny] = true // 如果可以走就标记并进一步dfs
				dfs(nx, ny, l+1, w)
				vis[nx][ny] = false // dfs完后回溯
			}
		}
	}
}

func newInit() {
	flag = false
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			vis[i][j] = false
		}
	}
}

func findWords(board [][]byte, words []string) []string {
	M = board
	fMap := make(map[byte][]string)
	for i, v := range board {
		for j, b := range v {
			if _, ok := fMap[b]; !ok {
				fMap[b] = make([]string, 0)
			}
			fMap[b] = append(fMap[b], fmt.Sprintf("%d,%d", i, j))
		}
	}

	res := make([]string, 0)
	for _, w := range words {
		if _, ok := fMap[w[0]]; !ok {
			continue
		}
		for _, idx := range fMap[w[0]] {
			str := strings.Split(idx, ",")
			x, _ := strconv.Atoi(str[0])
			y, _ := strconv.Atoi(str[1])
			newInit()
			vis[x][y] = true
			dfs(x, y, 0, w)
			if flag {
				res = append(res, w)
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findWords([][]byte{{'a', 'a'}}, []string{"aaa"}))
}
