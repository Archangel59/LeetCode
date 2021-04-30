package main

import (
	"fmt"
	"math"
)

func singleNumber(nums []int) int {
	// 把每个数按位存起来.
	var vis [40]int
	for _, v := range nums {
		x := int64(v + math.MaxUint32)
		idx := 1
		for x > 0 {
			if (x & 1) == 1 {
				vis[idx]++
			}
			x >>= 1
			idx++
		}
	}

	ans := 0
	for i := 1; i < 40; i++ {
		// 按位%3取出
		ans += (1 << (i - 1)) * (vis[i] % 3)
	}
	return ans - math.MaxUint32
}

func main() {
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}
