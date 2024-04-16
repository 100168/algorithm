package main

/*给定一个字符串 s ，找出 至多 包含 k 个不同字符的最长子串 T。*/

func lengthOfLongestSubstringKDistinct(s string, k int) int {

	hashSet := map[uint8]int{}

	left := 0
	right := 0

	ans := 0

	for right < len(s) {

		if len(hashSet) < k+1 {
			hashSet[s[right]]++
			right++
		}

		if len(hashSet) == k+1 {

			hashSet[s[left]]--
			if hashSet[s[left]] == 0 {
				delete(hashSet, s[left])
			}
			left++
		}

		if right-left > ans {
			ans = right - left
		}
	}

	return ans
}
