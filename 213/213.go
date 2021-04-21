package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// flag = 1 表示不包含第一个数
// flag = 0 表示包含第一个数
func getAns(nums []int, flag int) int {
	var num1 []int
	num1 = append(num1, nums...)

	// 最小值为第一个数
	ans := nums[1]

	for i := 3 + flag; i < len(nums)-1+flag; i++ {
		if i == 4 && flag == 1 {
			num1[i] = num1[i-2] + num1[i]
		} else {
			num1[i] = max(num1[i-3], num1[i-2]) + num1[i]
		}
	}

	for i := 1 + flag; i < len(nums)-1+flag; i++ {
		ans = max(ans, num1[i])
	}
	return ans
}

func rob(nums []int) int {
	nums = append([]int{0}, nums...)
	return max(getAns(nums, 0), getAns(nums, 1))
}

func main() {
	fmt.Println(rob([]int{1, 2, 1, 1}))
}
