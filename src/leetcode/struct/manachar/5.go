package main

/*
*

manachar

1. p 表示每个点为中心的最长回文长度
2. r表示回文半径又边界
3. c表示回文中心
*/
func longestPalindrome(s string) string {
	manaChar := "#"

	for _, v := range s {
		manaChar = manaChar + string(v) + "#"
	}
	n := len(manaChar)
	//每个点的最长回文半径
	p := make([]int, n)
	//c回文中心,r是回问半径（）
	c, r := 0, 0

	// a b c d
	maxLen := 0
	maxIndex := 0
	for i := 0; i < n; i++ {
		//自己也是回文串
		l := 1
		if r > i {
			l = min(r-i, p[2*c-i])
		}

		for i+l < n && i-l >= 0 && manaChar[i+l] == manaChar[i-l] {
			l++
		}

		if i+l > r {
			r = i + l
			c = i
		}
		maxLen = max(maxLen, l)
		if maxLen == l {
			maxIndex = i
		}
		p[i] = l
	}

	ans := ""

	if manaChar[maxIndex] != '#' {
		ans += string(manaChar[maxIndex])
		maxIndex++
		maxLen--
	}
	for i := maxIndex; i < maxIndex+maxLen; i++ {
		if manaChar[i] != '#' {
			ans += string(manaChar[i])
			ans = string(manaChar[i]) + ans
		}
	}
	return ans
}

func main() {

	println(longestPalindrome("babad"))
	println(longestPalindrome("aabbaa"))
}
