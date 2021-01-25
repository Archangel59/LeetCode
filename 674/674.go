package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	ans, l, n := 1, 1, len(nums)
	if n == 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		if nums[i-1] < nums[i] {
			l++
		} else {
			ans = max(ans, l)
			l = 1
		}
	}
	return max(ans, l)
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{}))
}
