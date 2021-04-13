package main

import (
	"sort"
	"strconv"
)

type List []string

func (a List) Len() int      { return len(a) }
func (a List) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a List) Less(i, j int) bool {
	n, m := a[i]+a[j], a[j]+a[i]
	return n > m
}

func largestNumber(nums []int) string {
	strList := List{}
	for _, v := range nums {
		strList = append(strList, strconv.Itoa(v))
	}
	sort.Sort(strList)
	res := ""
	flag := false
	for _, v := range strList {
		if v == "0" && !flag {
			continue
		}
		flag = true
		res += v
	}
	if len(res) == 0 {
		return "0"
	}
	return res
}

func main() {
	largestNumber([]int{3, 30, 34, 5, 9})
}
