package main

import (
	"fmt"
	"sort"
)

var mid, n int
var val [20]int
var v []int

func dfs(idx int) bool {
	if idx == len(v) {
		return true
	}
	for i := 0; i < n; i++ {
		if val[i]+v[idx] <= mid {
			val[i] += v[idx]
			if !dfs(idx + 1) {
				val[i] -= v[idx]
			} else {
				return true
			}
		}
		if val[i] == 0 || val[i]+v[idx] == mid {
			return false
		}
	}
	return false
}

func check() bool {
	// init
	for i := 0; i < n; i++ {
		val[i] = 0
	}
	return dfs(0)
}

func minimumTimeRequired(jobs []int, k int) int {
	// 倒序
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i] > jobs[j]
	})
	n, v = k, jobs

	// get l,r
	l, r := 0, 0
	for _, v := range jobs {
		r += v
	}
	l = r / k

	// 二分
	for l < r {
		mid = (l + r) >> 1
		if check() {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(minimumTimeRequired([]int{250, 250, 256, 251, 254, 254, 251, 255, 250, 252, 254, 255}, 10))
}
