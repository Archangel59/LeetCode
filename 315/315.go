package main

import (
	"fmt"
	"sort"
)

var Tree [400010]struct {
	l, r, value int
}

// 离散化
func getHashArr(nums []int) []int {
	type node struct {
		value, idx int
	}
	type List []node
	data := List{}
	for k, v := range nums {
		data = append(data, node{value: v, idx: k})
	}
	sort.Slice(data, func(i, j int) bool {
		if data[i].value == data[j].value {
			return data[i].idx < data[j].idx
		}
		return data[i].value < data[j].value
	})

	head := data[0].value
	idx := 1
	for i := 0; i < len(data); i++ {
		if i == 0 {
			data[i].value = 1
			continue
		}
		if data[i].value != head {
			idx++
		}
		data[i].value = idx
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].idx < data[j].idx
	})

	arr := make([]int, len(data))
	for k, v := range data {
		arr[k] = v.value
	}

	return arr
}

// build线段树
func build(l, r, k int) {
	Tree[k].l = l
	Tree[k].r = r
	if l == r {
		Tree[k].value = 0
		return
	}
	mid := (l + r) >> 1
	build(l, mid, k<<1)
	build(mid+1, r, (k<<1)+1)
	Tree[k].value = Tree[k<<1].value + Tree[(k<<1)+1].value
}

// 区间获取sum
func getRange(l, r, k int) int {
	if Tree[k].l >= l && Tree[k].r <= r {
		return Tree[k].value
	}
	mid := (Tree[k].l + Tree[k].r) >> 1
	ans := 0
	if l <= mid {
		ans += getRange(l, r, k<<1)
	}
	if mid < r {
		ans += getRange(l, r, (k<<1)+1)
	}
	return ans
}

// 更新一个
func updateOne(idx, k int) {
	if Tree[k].l == Tree[k].r {
		Tree[k].value++
		return
	}
	mid := (Tree[k].l + Tree[k].r) >> 1
	if idx <= mid {
		updateOne(idx, k<<1)
	} else {
		updateOne(idx, (k<<1)+1)
	}
	Tree[k].value = Tree[k<<1].value + Tree[(k<<1)+1].value
}

func countSmaller(nums []int) []int {
	arr := getHashArr(nums)
	build(1, len(arr), 1)
	res := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		updateOne(arr[i], 1)
		if arr[i] == 1 {
			res[i] = 0
			continue
		}
		res[i] = getRange(1, arr[i]-1, 1)
	}

	return res
}

func main() {
	fmt.Println(countSmaller([]int{1, 9, 7, 8, 5}))
}
