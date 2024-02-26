package main

func lengthOfLongestSubstring(s string) int {
	l := 0
	r := 0
	sMap := make(map[uint8]bool)

	ans := 0

	for ; r < len(s); r++ {
		for ; sMap[s[r]]; l++ {
			sMap[s[l]] = false
		}
		sMap[s[r]] = true
		ans = max(ans, r-l+1)
	}
	return ans
}
