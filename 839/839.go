package main

import "fmt"

var fa = [1100]int{}

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

func check(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	n := len(a)
	var dis []int

	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			dis = append(dis, i)
		}
	}
	switch len(dis) {
	case 0:
		return true
	case 2:
		if a[dis[0]] == b[dis[1]] && a[dis[1]] == b[dis[0]] {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

func numSimilarGroups(strs []string) int {
	n := len(strs)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if check(strs[i], strs[j]) {
				union(i, j)
			}
		}
	}
	mp := make(map[int]bool)
	for i := 0; i < n; i++ {
		mp[find(fa[i])] = true
	}
	return len(mp)
}

func main() {
	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star"}))
}
