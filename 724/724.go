package main

import "fmt"

func pivotIndex(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	a, b := make([]int, n), make([]int, n)
	a[0], b[n-1] = nums[0], nums[n-1]
	for i, j := 1, n-2; i < n; i, j = i+1, j-1 {
		a[i] = a[i-1] + nums[i]
		b[j] = b[j+1] + nums[j]
	}
	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
