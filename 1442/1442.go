package main

import "fmt"

type node struct {
	l, r, val int
}

var tree [305 * 4]node
var a []int

func build(l, r, k int, idx *int) {
	tree[k].l, tree[k].r = l, r
	if l == r {
		tree[k].val = a[*idx]
		*idx++
		return
	}
	mid := (l + r) >> 1
	build(l, mid, k<<1, idx)
	build(mid+1, r, (k<<1)+1, idx)
	tree[k].val = tree[k<<1].val ^ tree[(k<<1)+1].val
}

func getRange(l, r, k int) int {
	if tree[k].l >= l && tree[k].r <= r {
		return tree[k].val
	}
	ans, mid := 0, (tree[k].l+tree[k].r)>>1
	if l <= mid {
		ans ^= getRange(l, r, k<<1)
	}
	if mid < r {
		ans ^= getRange(l, r, (k<<1)+1)
	}
	return ans
}

func countTriplets(arr []int) int {
	a = arr
	idx, ans := 0, 0
	build(1, len(a), 1, &idx)
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if getRange(i+1, j+1, 1) == 0 {
				ans += j - i
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(countTriplets([]int{2, 3, 1, 6, 7}))
}
