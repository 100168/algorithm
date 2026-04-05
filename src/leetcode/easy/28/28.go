package main

import "fmt"

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	if m == 0 {
		return 0
	}
	i := 1
	j := 0
	next := make([]int, m)
	for i < m {
		if needle[i] == needle[j] {
			next[i] = j + 1
			i++
			j++
		} else {
			if j > 0 {
				j = next[j-1]
			} else {
				i++
			}
		}

	}
	r := 0
	l := 0
	for r < n {
		if haystack[r] == needle[l] {
			r++
			if l == m-1 {
				return r - m
			}
			l++
		} else {
			if l > 0 {
				l = next[l-1]
			} else {
				r++
			}
		}
	}
	return -1
}
func main() {
	str := strStr("helloooolllolllooolllooooo",
		"llooooo")
	fmt.Println(str)
}
