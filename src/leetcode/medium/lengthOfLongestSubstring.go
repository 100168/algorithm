package main

import "fmt"

/*给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。*/
func lengthOfLongestSubstring(s string) int {
	longest := 0
	left := 0
	right := 0
	isInclude := map[byte]bool{}
	n := len(s)

	for right < n {
		if !isInclude[s[right]] {
			isInclude[s[right]] = true
			right++
		} else {
			if right-left > longest {
				longest = right - left
			}
			isInclude[s[left]] = false
			left++
		}
	}
	if right-left > longest {
		longest = right - left
	}
	return longest
}
func main() {
	fmt.Println("hhh")
}
