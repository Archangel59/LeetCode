package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	nums = append([]int{0}, nums...)
	for i := 3; i < len(nums); i++ {
		nums[i] = max(nums[i-3], nums[i-2]) + nums[i]
	}
	ans := 0
	for i := 1; i < len(nums); i++ {
		ans = max(ans, nums[i])
	}
	return ans
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
