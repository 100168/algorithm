package main

/*
给你两个字符串 s1 和 s2 ，两个字符串长度都为 n ，且只包含 小写 英文字母。
你可以对两个字符串中的 任意一个 执行以下操作 任意 次：
选择两个下标 i 和 j ，满足 i < j 且 j - i 是 偶数，然后 交换 这个字符串中两个下标对应的字符。
如果你可以让字符串 s1 和 s2 相等，那么返回 true ，否则返回 false 。
*/
func checkStrings(s1 string, s2 string) bool {

	cntS1 := make([]int, 26)
	cntS2 := make([]int, 26)

	cntS11 := make([][]int, 26)
	cntS22 := make([][]int, 26)

	for i := range cntS11 {
		cntS11[i] = make([]int, 2)
	}
	for i := range cntS22 {
		cntS22[i] = make([]int, 2)
	}
	if len(s1) != len(s2) {
		return false
	}

	for _, i2 := range s1 {
		cur := int(i2 - 'a')
		cntS1[cur]++
	}

	for i, i2 := range s2 {
		cur := int(i2 - 'a')
		cntS2[cur]++
		if i&1 == 0 {
			cntS22[cur][0]++
		} else {
			cntS22[cur][1]++
		}
	}
	for i, i2 := range s1 {
		cur := int(i2 - 'a')
		if cntS1[cur] != cntS2[cur] {
			return false
		}
		cntS22[cur][i&1]--
		if cntS22[cur][i&1] < 0 {
			return false
		}
	}
	return true

}

func main() {

	s1 := "jbkbprwoteofzxhegjdlqhxyxwukzvxppcapcogguvextcqyeqalcfypllwcbcvhwjokuqhozdjpjggazaegydmbkvlkerrwwwq"
	s2 := "zkjothlxdbtlpvgcbcgbqczdjkzpkjlwvgupmgalwazkqqqfocwqebuywhovclyphyarydewgewerrjkpgxoxxhoxpcfwvuaeje"
	println(checkStrings(s1, s2))

}
