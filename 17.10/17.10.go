package main

import "fmt"

func majorityElement(nums []int) int {
	val, cnt := 0, 0
	for _, v := range nums {
		if cnt == 0 {
			val = v
		}
		if val == v {
			cnt++
		} else {
			cnt--
		}
	}
	cnt = 0
	for _, v := range nums {
		if v == val {
			cnt++
		}
	}
	if cnt > len(nums)/2 {
		return val
	}
	return -1
}

func main() {
	fmt.Println(majorityElement([]int{1, 2}))
}
