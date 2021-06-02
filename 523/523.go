package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	mp := make(map[int]int)
	mp[0] = -1
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			nums[i] += nums[i-1]
		}
		nums[i] %= k
		if _, ok := mp[nums[i]]; ok {
			if i-mp[nums[i]] > 1 {
				return true
			}
		} else {
			mp[nums[i]] = i
		}
	}
	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{5, 0, 0, 0}, 3))
}
