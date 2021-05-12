package main

import "fmt"

const MaxN = 30010

type node struct {
	l, r, v int
}

var Tree = [MaxN * 4]node{}
var a []int

func build(l, r, k int, idx *int) {
	Tree[k].l, Tree[k].r = l, r
	if l == r {
		Tree[k].v = a[*idx]
		*idx++
		return
	}
	mid := (l + r) >> 1
	build(l, mid, k<<1, idx)
	build(mid+1, r, (k<<1)+1, idx)
	Tree[k].v = Tree[k<<1].v ^ Tree[(k<<1)+1].v
}

func getRange(l, r, k int) int {
	if Tree[k].l >= l && Tree[k].r <= r {
		return Tree[k].v
	}
	mid, ans := (Tree[k].l+Tree[k].r)>>1, 0
	if l <= mid {
		ans ^= getRange(l, r, k<<1)
	}
	if mid < r {
		ans ^= getRange(l, r, (k<<1)+1)
	}
	return ans
}

func xorQueries(arr []int, queries [][]int) []int {
	a = arr
	idx := 0
	build(1, len(arr), 1, &idx)
	ans := make([]int, len(queries))
	for k, v := range queries {
		ans[k] = getRange(v[0]+1, v[1]+1, 1)
	}
	return ans
}

func main() {
	//[4,8,2,10]
	//[[2,3],[1,3],[0,0],[0,3]]
	fmt.Println(xorQueries([]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}))
}
