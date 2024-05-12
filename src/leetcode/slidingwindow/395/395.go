package main

func longestSubstring(s string, k int) int {

	n := len(s)
	ans := 0
	for i := 1; i <= 26; i++ {
		cnt := make(map[int]int)
		kind := make(map[int]bool)

		l := 0
	next:
		for r := 0; r < n; r++ {
			cur := int(s[r] - 'a')
			kind[cur] = true

			for len(kind) > i {
				left := int(s[l] - 'a')
				cnt[left]--
				if cnt[left] == 0 {
					delete(cnt, left)
					delete(kind, left)
				}
				l++
			}
			for _, v := range cnt {
				if v < k {
					continue next
				}
			}
			ans = max(ans, r-l+1)
		}

	}

	return ans

}
