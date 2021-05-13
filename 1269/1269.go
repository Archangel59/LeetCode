package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numWays(steps int, arrLen int) int {
	dp := make([][]int, steps+1)
	maxColumn := min(arrLen-1, steps)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, maxColumn+1)
	}
	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j <= maxColumn; j++ {
			dp[i][j] = dp[i-1][j]
			if j > 0 {
				dp[i][j] += dp[i-1][j-1]
			}
			if j < maxColumn {
				dp[i][j] += dp[i-1][j+1]
			}
			dp[i][j] %= 1e9 + 7
		}
	}

	return dp[steps][0]
}

func main() {
	fmt.Println(numWays(3, 2))
}
