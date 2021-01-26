package main

import "fmt"

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func numEquivDominoPairs(dominoes [][]int) int {
	var a [100]int
	ans := 0
	for _, v := range dominoes {
		x := min(v[0]*10+v[1], v[1]*10+v[0])
		ans += a[x]
		a[x]++
	}
	return ans
}

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
}
