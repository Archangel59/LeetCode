package main

import "fmt"

func cal(val, v int) (a int, b int) {
	if val-v < 0 {
		return v - val, v + val
	}
	return val - v, v + val
}

func lastStoneWeightII(stones []int) int {
	a, b := make(map[int]int), make(map[int]int)
	a[stones[0]] = 1
	// dp[i][j] = dp[i-1][j-stones] + dp[i-1][j+stones]
	for k, v := range stones {
		if k == 0 {
			continue
		}
		for val, _ := range a {
			c, d := cal(val, v)
			b[c] = 1
			b[d] = 1
		}
		a = b
		b = make(map[int]int)
	}
	max := 3000
	for k, _ := range a {
		if k < max {
			max = k
		}
	}
	return max
}

func main() {
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}
