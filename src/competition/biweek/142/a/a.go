package main

/*
*
Alice 正在她的电脑上输入一个字符串。但是她打字技术比较笨拙，她 可能 在一个按键上按太久，导致一个字符被输入 多次 。

尽管 Alice 尽可能集中注意力，她仍然可能会犯错 至多 一次。

给你一个字符串 word ，它表示 最终 显示在 Alice 显示屏上的结果。

请你返回 Alice 一开始可能想要输入字符串的总方案数。

示例 1：

输入：word = "abbcccc"

输出：5

解释：

可能的字符串包括："abbcccc" ，"abbccc" ，"abbcc" ，"abbc" 和 "abcccc" 。
*/
func possibleStringCount(word string) int {

	ans := 1
	for i, v := range word[1:] {
		if byte(v) == word[i] {
			ans++
		}
	}
	return ans
}
