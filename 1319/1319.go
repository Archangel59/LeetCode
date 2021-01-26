package main

import "fmt"

var fa = [100010]int{}

func find(x int) int {
	if x != fa[x] {
		fa[x] = find(fa[x])
	}
	return fa[x]
}

func union(x, y int) {
	if find(x) != find(y) {
		fa[find(x)] = fa[find(y)]
	}
}

func makeConnected(n int, connections [][]int) int {
	if n > len(connections)+1 {
		return -1
	}
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	for _, v := range connections {
		union(v[0], v[1])
	}
	mp := make(map[int]bool)
	for i := 0; i < n; i++ {
		mp[find(fa[i])] = true
	}

	return len(mp) - 1
}

func main() {
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}))
}
