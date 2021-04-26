package main

import (
	"fmt"
	"math"
)

func check(arr []int, D, size int) bool {
	ans, num := 0, 0
	for _, v := range arr {
		// 存在某件物品大于当前 size
		if v > size {
			return false
		}
		num += v
		if num > size {
			ans++
			num = v
		}
	}
	if num > 0 {
		ans++
	}
	return ans <= D
}

func shipWithinDays(weights []int, D int) int {
	l, r := 0, math.MaxInt32
	for l < r {
		mid := (l + r) >> 1
		if check(weights, D, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {

	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
}
