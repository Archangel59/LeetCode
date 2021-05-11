package main

import "fmt"

func decode(encoded []int) []int {
	n := len(encoded) + 1
	sum, head := 0, 0
	for i := 1; i <= n; i++ {
		sum ^= i
	}
	for i := 1; i < n; i += 2 {
		head ^= encoded[i]
	}

	ans := make([]int, n)
	ans[0] = head ^ sum
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] ^ encoded[i-1]
	}
	return ans
}

func main() {
	fmt.Println(decode([]int{6, 5, 4, 6}))
}

/*
输入：encoded = [6,5,4,6]
		输出：[2,4,1,5,3]
*/
