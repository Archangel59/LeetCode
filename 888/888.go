package main

import "fmt"

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0
	ans := make(map[int]int)
	for _, v := range A {
		sumA += v
	}
	for _, v := range A {
		ans[sumA-v-v] = v
	}
	for _, v := range B {
		sumB += v
	}
	for _, v := range B {
		if a, ok := ans[sumB-v-v]; ok {
			return []int{a, v}
		}
	}
	return nil
}

func main() {
	fmt.Println(fairCandySwap([]int{1, 2, 5}, []int{2, 4}))
}
