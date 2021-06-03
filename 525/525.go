package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxLength(nums []int) int {
	ans, x, y := 0, 0, 0
	mp := make(map[int]map[int]int)
	mp[0] = make(map[int]int)
	mp[0][0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			x++
		} else {
			y++
		}
		m := min(x, y)
		x -= m
		y -= m
		if idx, ok := mp[x][y]; ok {
			ans = max(ans, i-idx)
		} else {
			if _, ok = mp[x]; !ok {
				mp[x] = make(map[int]int)
			}
			mp[x][y] = i
		}
	}

	return ans
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1, 0}))
}
