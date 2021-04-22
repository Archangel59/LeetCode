package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func rmqInit(arr []int) [][]int {
	n := len(arr)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 20)
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = arr[i-1]
	}
	for j := 1; (1 << j) <= n; j++ {
		for i := 1; i+(1<<(j-1)) <= n; i++ {
			dp[i][j] = dp[i][j-1] & dp[i+(1<<(j-1))][j-1]
		}
	}

	return dp
}

func getRange(dp [][]int, l, r int) int {
	k := int(math.Log2(float64(r - l + 1)))
	return dp[l][k] & dp[r-(1<<k)+1][k]
}

// 二分查找
func solve(dp [][]int, r, target int) int {
	ans := math.MaxInt32
	left := 1
	right := r
	for left <= right {
		mid := (left + right) >> 1
		midAns := getRange(dp, mid, r)
		if midAns == target {
			return 0
		}
		if midAns > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right == 0 {
		ans = min(ans, getRange(dp, left, r)-target)
	} else if left == r+1 {
		ans = min(ans, target-getRange(dp, right, r))
	} else {
		ans = min(ans, min(getRange(dp, left, r)-target, target-getRange(dp, right, r)))
	}
	return ans
}

func closestToTarget(arr []int, target int) int {
	dp := rmqInit(arr)
	ans := abs(arr[0] - target)
	for r := 2; r <= len(arr); r++ {
		ans = min(ans, solve(dp, r, target))
	}

	return ans
}

func main() {
	fmt.Println(closestToTarget([]int{10, 9, 11, 22, 32}, 4))
}
