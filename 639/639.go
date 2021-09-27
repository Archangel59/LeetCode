package main

import "fmt"

func numDecodings(s string) int {
	a, b, c, mod := 0, 1, 0, 1000000007
	for i := 0; i < len(s); i++ {
		one := func() int {
			switch s[i] {
			case '*':
				return 9
			case '0':
				return 0
			default:
				return 1
			}
		}

		two := func() int {
			if s[i] == '*' && s[i-1] == '*' {
				return 15
			}
			if s[i-1] == '*' {
				if s[i] >= '0' && s[i] <= '6' {
					return 2
				} else {
					return 1
				}
			}
			if s[i] == '*' {
				if s[i-1] == '1' {
					return 9
				}
				if s[i-1] == '2' {
					return 6
				}
				return 0
			}
			if s[i-1] != '0' && (s[i-1]-'0')*10+(s[i]-'0') <= 26 {
				return 1
			}
			return 0
		}

		c = b * one() % mod
		if i > 0 {
			c += a * two() % mod
			c %= mod
		}
		a, b = b, c
	}
	return c
}

func main() {
	fmt.Println(numDecodings("2*"))
}
