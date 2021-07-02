package main

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	ans := 0
	for _, v := range costs {
		if coins >= v {
			ans++
			coins -= v
		}
	}
	return ans
}

func main() {
	fmt.Println(maxIceCream([]int{1, 6, 3, 1, 2, 5}, 20))
}
