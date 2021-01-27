package main

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	mp := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		num := nums[i] / (t + 1)
		if nums[i] < 0 {
			num--
		}
		// 已经存在
		if _, ok := mp[num]; ok {
			return true
		}
		// 上一个桶存在
		if v, ok := mp[num-1]; ok && nums[i]-v <= t {
			return true
		}
		// 下一个桶存在
		if v, ok := mp[num+1]; ok && v-nums[i] <= t {
			return true
		}
		mp[num] = nums[i]
		if i >= k {
			delete(mp, nums[i-k]/(t+1))
		}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{2147483647, -1, 2147483647}, 1, 2147483647))
}
