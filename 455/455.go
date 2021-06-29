package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans, idx := 0, 0
	for _, v := range g {
		for i := idx; i < len(s); i++ {
			if s[i] >= v {
				idx = i + 1
				ans++
				break
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
}
