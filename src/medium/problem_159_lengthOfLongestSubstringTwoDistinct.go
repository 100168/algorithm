package main

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {

	left := 0
	right := 0

	hashSet := map[uint8]int{}

	ans := 0
	for right < len(s) {
		if len(hashSet) < 3 {
			hashSet[s[right]]++
			right++
		}
		if len(hashSet) == 3 {
			hashSet[s[left]]--
			if hashSet[s[left]] == 0 {
				delete(hashSet, s[left])
			}
			left++
		}
		if ans < right-left+1 {
			ans = right - left
		}
	}
	return ans
}

func main() {
	distinct := lengthOfLongestSubstringTwoDistinct("bacbbcabcabbcc")

	fmt.Println(distinct)
}
