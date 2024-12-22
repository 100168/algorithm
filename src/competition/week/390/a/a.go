package main

/*
*
滑动窗口
*/
func maximumLengthSubstring(s string) int {

	cnt := make(map[uint8]int)

	l := 0
	ans := 0
	n := len(s)
	for i := 0; i < n; i++ {
		cur := s[i] - 'a'
		cnt[cur]++
		for cnt[cur] > 2 {
			cnt[s[l]-'a']--
			l++
		}
		ans = max(ans, i-l+1)
	}
	return ans
}
