package main

import (
	"fmt"
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	ans := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > ans {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumElementAfterDecrementingAndRearranging([]int{1, 5, 2, 3, 4}))
}
