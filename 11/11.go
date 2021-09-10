package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	l, r, ans := 0, len(height)-1, 0
	for {
		if l >= r {
			break
		}
		fmt.Println(ans, l, r)
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}
	return ans
}

func main() {
	fmt.Println(maxArea([]int{1, 2, 1}))
}
