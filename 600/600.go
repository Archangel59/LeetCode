package main

import (
	"fmt"
)

func findIntegers(n int) int {
	// 初始化斐波拉契数列
	var f [34]int
	f[0] = 1
	f[1] = 2
	for i := 2; i <= 32; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	ans := 0
	for i := 32; i >= 0; i-- {
		if n&(1<<i) == 0 { // 按位去与, 如果为0则表示n该位为0
			continue
		}
		ans += f[i]
		if n&(1<<(i+1)) != 0 { // 如果有连续的1, 表明不合法, 后续不用进行判断
			return ans
		}
	}
	return ans + 1
}

func main() {
	fmt.Println(findIntegers(13))
}
