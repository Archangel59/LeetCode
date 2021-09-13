package main

import "fmt"

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for _, p := range points {
		mp := make(map[int]int)
		for _, q := range points {
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			mp[dis]++
		}
		for _, v := range mp {
			ans += v * (v - 1)
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfBoomerangs([][]int{{1, 1}, {2, 2}, {3, 3}}))
}
