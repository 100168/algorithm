package main

func strStr(haystack string, needle string) int {

	m := len(haystack)
	n := len(needle)

	next := make([]int, n)

	for i := 1; i < n; i++ {
		j := next[i-1]
		for j != 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[j] == needle[i] {
			next[i] = j + 1
		}
	}

	j := 0
	for i := 0; i < m && j < n; i++ {
		for j != 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}

		if j == len(next) {
			return i - j + 1
		}

	}
	return -1
}

func strStr2(haystack string, needle string) int {

	p1 := int64(131313)
	m := len(haystack)
	n := len(needle)
	pm := make([]int64, m+1)
	hm := make([]int64, m+1)
	hn := int64(0)
	pm[0] = 1
	for i := 1; i <= m; i++ {
		hm[i] = hm[i-1]*p1 + int64(haystack[i-1]-'a')
		pm[i] = pm[i-1] * p1
	}
	for i := 1; i <= n; i++ {
		hn = hn*p1 + int64(needle[i-1]-'a')
	}
	for end := n; end <= m; end++ {
		start := end - n
		cur := hm[end] - hm[start]*pm[n]
		if cur == hn {
			return start
		}
	}
	return -1
}

func main() {

	println(strStr("butsad", "sad"))
	println(strStr2("butsad", "sad"))
}
