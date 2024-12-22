package main

/**

给你一个字符串 target。

Alice 将会使用一种特殊的键盘在她的电脑上输入 target，这个键盘 只有两个 按键：

按键 1：在屏幕上的字符串后追加字符 'a'。
按键 2：将屏幕上字符串的 最后一个 字符更改为英文字母表中的 下一个 字符。例如，'c' 变为 'd'，'z' 变为 'a'。
注意，最初屏幕上是一个空字符串 ""，所以她 只能 按按键 1。

请你考虑按键次数 最少 的情况，按字符串出现顺序，返回 Alice 输入 target 时屏幕上出现的所有字符串列表。



示例 1：

输入： target = "abc"

输出： ["a","aa","ab","aba","abb","abc"]

解释：

Alice 按键的顺序如下：

按下按键 1，屏幕上的字符串变为 "a"。
按下按键 1，屏幕上的字符串变为 "aa"。
按下按键 2，屏幕上的字符串变为 "ab"。
按下按键 1，屏幕上的字符串变为 "aba"。
按下按键 2，屏幕上的字符串变为 "abb"。
按下按键 2，屏幕上的字符串变为 "abc"。

*/

func stringSequence(target string) []string {

	ans := []string{""}

	for _, v := range target {
		ans = append(ans, ans[len(ans)-1]+"a")
		for j := 'b'; j <= v; j++ {

			last := ans[len(ans)-1]
			last = last[:len(last)-1]
			ans = append(ans, last+string(j))
		}
	}
	return ans[1:]
}
