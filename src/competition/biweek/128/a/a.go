package a

func scoreOfString(s string) int {

	ans := 0
	for i := 1; i < len(s); i++ {
		cur := int(s[i]-'a') - int(s[i-1]-'a')
		ans += abs(cur)
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
