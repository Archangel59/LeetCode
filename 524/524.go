package main

import (
	"fmt"
	"sort"
)

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) == len(dictionary[j]) {
			return dictionary[i] < dictionary[j]
		}
		return len(dictionary[i]) > len(dictionary[j])
	})
	for _, words := range dictionary {
		check := func() bool {
			idx, cnt := 0, 0
			for wi := range words {
				for i := idx; i < len(s); i++ {
					if s[i] == words[wi] {
						cnt++
						idx = i + 1
						break
					}
				}
			}
			return cnt == len(words)
		}

		if check() {
			return words
		}
	}
	return ""
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea", "abpcplaaa", "abpcllllll", "abccclllpppeeaaaa"}))
}
