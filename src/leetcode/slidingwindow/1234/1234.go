package main

/*
*
有一个只含有 'Q', 'W', 'E', 'R' 四种字符，且长度为 n 的字符串。
假如在该字符串中，这四个字符都恰好出现 n/4 次，那么它就是一个「平衡字符串」。
给你一个这样的字符串 s，请通过「替换一个子串」的方式，使原字符串 s 变成一个「平衡字符串」。
你可以用和「待替换子串」长度相同的 任何 其他字符串来完成替换。
请返回待替换子串的最小可能长度。
如果原字符串自身就是一个平衡字符串，则返回 0。
*/

// 二分+定长划窗
func balancedString(s string) int {

	n := len(s)
	cnt := make(map[uint8]int)
	for i := range s {
		cnt[s[i]]++
	}
	isB := true
	for _, v := range cnt {
		if v != n/4 {
			isB = false
		}
	}

	//EEWWWWEE
	if isB {
		return 0
	}

	l, r := 0, n-1

	for l <= r {
		m := (r + l) / 2
		if check(s, m, cnt) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l

}

func check(s string, m int, c map[uint8]int) bool {

	n := len(s)
	w := make(map[uint8]int)
	l := 0
next:
	for r := 0; r < len(s); r++ {
		w[s[r]]++
		if r-l+1 > m {
			w[s[l]]--
			l++
		}
		for k, v := range c {
			if v-w[k] > n/4 {
				continue next
			}
		}
		return true
	}
	return false

}

// 直接滑窗
func balancedString2(s string) int {

	n := len(s)
	cnt := make(map[uint8]int)
	for i := range s {
		cnt[s[i]]++
	}
	isB := true
	for _, v := range cnt {
		if v != n/4 {
			isB = false
		}
	}

	//EEWWWWEE
	if isB {
		return 0
	}

	ans := n
	w := make(map[uint8]int)
	l := 0
next:
	for r := 0; r < len(s); r++ {
		w[s[r]]++
		for l <= r && cnt[s[l]]-w[s[l]] < n/4 {
			w[s[l]]--
			l++
		}
		for k, v := range cnt {
			if v-w[k] > n/4 {
				continue next
			}
		}
		ans = min(ans, r-l+1)
	}
	return ans

}

func main() {
	println(balancedString("EEWWWWEE"))
	println(balancedString2("QQQW"))
}
