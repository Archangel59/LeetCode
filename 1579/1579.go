package main

import "fmt"

var fa = [100010]int{}
var fb = [100010]int{}

func findFa(x int) int {
	if x != fa[x] {
		fa[x] = findFa(fa[x])
	}
	return fa[x]
}

func findFb(x int) int {
	if x != fb[x] {
		fb[x] = findFb(fb[x])
	}
	return fb[x]
}

func getThreeArr(edges [][]int) (a, b, c [][]int) {
	for _, v := range edges {
		if v[0] == 1 {
			a = append(a, v)
		} else if v[0] == 2 {
			b = append(b, v)
		} else {
			c = append(c, v)
		}
	}
	return
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		fa[i] = i
	}
	a, b, c := getThreeArr(edges)
	// 先处理公共边
	for _, v := range c {
		if findFa(v[1]) != findFa(v[2]) {
			fa[findFa(v[1])] = fa[findFa(v[2])]
		} else {
			ans++
		}
	}
	// 复制一份给B使用
	fb = fa
	// 接着处理A
	for _, v := range a {
		if findFa(v[1]) != findFa(v[2]) {
			fa[findFa(v[1])] = fa[findFa(v[2])]
		} else {
			ans++
		}
	}
	// 最后处理B
	for _, v := range b {
		if findFb(v[1]) != findFb(v[2]) {
			fb[findFb(v[1])] = fb[findFb(v[2])]
		} else {
			ans++
		}
	}
	// 判断fa和fb是否是完整图
	for i := 2; i <= n; i++ {
		if findFa(1) != findFa(i) {
			return -1
		}
		if findFb(1) != findFb(i) {
			return -1
		}
	}

	return ans
}

func main() {
	fmt.Println(maxNumEdgesToRemove(4, [][]int{{3, 2, 3}, {1, 1, 2}, {2, 3, 4}}))
}
