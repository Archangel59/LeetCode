package main

import "fmt"

func cal(x, y int) []int {
	if y == 0 {
		return []int{x}
	}
	res := make([]int, 0)
	z := x % y
	b := x / y
	for i := 0; i < z; i++ {
		res = append(res, b+1)
	}
	for i := 0; i < y-z; i++ {
		res = append(res, b)
	}
	return res
}

func fullJustify(words []string, maxWidth int) []string {
	res := make([][]string, 0)
	now := make([]string, 0)
	num, idx := 0, 0
	for _, w := range words {
		if num > 0 {
			num++
		}
		if num+len(w) > maxWidth {
			idx++
			res = append(res, now)
			now = make([]string, 0)
			num = 0
		}
		now = append(now, w)
		num += len(w)
	}
	if len(now) > 0 {
		res = append(res, now)
	}

	ans := make([]string, 0)
	for k, d := range res {
		if k == len(res)-1 {
			break
		}
		tot := 0
		for _, w := range d {
			tot += len(w)
		}
		b := cal(maxWidth-tot, len(d)-1)
		line := d[0]
		for j, v := range b {
			for i := 1; i <= v; i++ {
				line += " "
			}
			if j+1 < len(d) {
				line += d[j+1]
			}
		}
		ans = append(ans, line)
	}

	line := ""
	for _, v := range res[len(res)-1] {
		line += v
		if len(line) < maxWidth {
			line += " "
		}
	}

	for i := len(line); i < maxWidth; i++ {
		line += " "
	}

	return append(ans, line)
}

func main() {
	fmt.Println(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
}
