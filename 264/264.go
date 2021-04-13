package main

import "fmt"

func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func nthUglyNumber(n int) int {
	dp := make([]int, 2000)
	dp[1] = 1
	x, y, z := 1, 1, 1
	for i := 2; i <= n; i++ {
		X, Y, Z := dp[x]*2, dp[y]*3, dp[z]*5
		min := getMin(X, getMin(Y, Z))
		dp[i] = min
		if min == X {
			x++
		}
		if min == Y {
			y++
		}
		if min == Z {
			z++
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
