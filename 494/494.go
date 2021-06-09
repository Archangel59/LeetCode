package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findTargetSumWays(nums []int, target int) int {
	n, sum := len(nums), 0
	a := [40010]int{}
	b := [40010]int{}
	for _, v := range nums {
		sum += abs(v)
	}
	a[sum] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= sum+sum; j++ {
			b[j] = a[j+nums[i-1]]
			if j-nums[i-1] >= 0 {
				b[j] += a[j-nums[i-1]]
			}
		}
		a = b
	}

	return a[target+sum]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

// f[i][j] 代表考虑前 i 个数，当前计算结果为 j 的方案数，令 nums 下标从 1 开始
// f[i][j] = f[i-1][j-1] + f[i-1][j+1]
