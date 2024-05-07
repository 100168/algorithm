package main

/*
*
给你一个由字符 'a'、'b'、'c' 组成的字符串 s 和一个非负整数 k 。每分钟，你可以选择取走 s 最左侧 还是 最右侧 的那个字符。

你必须取走每种字符 至少 k 个，返回需要的 最少 分钟数；如果无法取到，则返回 -1 。
*/
func takeCharacters(s string, k int) int {

	//cnt[0]-k>=y
	n := len(s)
	cnt := make([]int, 3)

	for i := 0; i < n; i++ {
		cur := int(s[i] - 'a')
		cnt[cur]++
	}

	w := make([]int, 3)

	for _, v := range cnt {
		if v < k {
			return -1
		}
	}
	l := 0
	ans := -1
	for r := 0; r < n; r++ {
		cur := int(s[r] - 'a')
		w[cur]++
		for w[cur] > cnt[cur]-k {
			left := int(s[l] - 'a')
			w[left]--
			l++
		}
		ans = max(ans, r-l+1)

	}

	return n - ans

}
