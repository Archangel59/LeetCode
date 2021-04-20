package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	// get next
	next := make([]int, len(needle)+1)
	i, j := -1, 0
	next[0] = -1
	for j < len(needle) {
		if i == -1 || needle[i] == needle[j] {
			i, j = i+1, j+1
			next[j] = i
		} else {
			i = next[i]
		}
	}
	// kmp
	i, j = 0, 0
	for j < len(haystack) {
		if i == -1 || needle[i] == haystack[j] {
			i, j = i+1, j+1
		} else {
			i = next[i]
		}
		if i == len(needle) {
			return j - i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("aa", "aa"))
}
