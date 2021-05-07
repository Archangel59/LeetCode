package main

import "fmt"

func xorOperation(n int, start int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans ^= start + i + i
	}
	return ans
}

func main() {
	fmt.Println(xorOperation(5, 0))
}
