package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	a := make([]int, 0)
	a = append(a, nums1...)
	sort.Ints(a)

	ans := 0
	for k := range nums1 {
		l, r := 0, len(a)-1
		for {
			if r-l <= 1 {
				break
			}
			mid := (l + r) >> 1
			if a[mid] > nums2[k] {
				r = mid
			} else if a[mid] < nums2[k] {
				l = mid
			} else {
				l = mid
				break
			}
		}
		old, ll, rr := abs(nums1[k]-nums2[k]), abs(a[l]-nums2[k]), abs(a[r]-nums2[k])
		if old-ll > old-rr {
			if ans < old-ll {
				ans = old - ll
			}
		} else {
			if ans < old-rr {
				ans = old - rr
			}
		}
	}

	sum, mod := 0, 1000000007
	for k := range nums1 {
		sum += abs(nums1[k] - nums2[k])
		sum %= mod
	}

	return (sum + mod - ans) % mod
}

func main() {
	fmt.Println(minAbsoluteSumDiff([]int{1, 7, 5}, []int{2, 3, 5}))
}
