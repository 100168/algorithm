package main

/*
给你一个小写英文字母组成的字符串 s 和一个二维整数数组 shifts ，其中 shifts[i] = [starti, endi, directioni] 。对于每个 i ，将 s 中从下标 starti 到下标 endi （两者都包含）所有字符都进行移位运算，如果 directioni = 1 将字符向后移位，如果 directioni = 0 将字符向前移位。

将一个字符 向后 移位的意思是将这个字符用字母表中 下一个 字母替换（字母表视为环绕的，所以 'z' 变成 'a'）。类似的，将一个字符 向前 移位的意思是将这个字符用字母表中 前一个 字母替换（字母表是环绕的，所以 'a' 变成 'z' ）。

请你返回对 s 进行所有移位操作以后得到的最终字符串。

示例 1：

输入：s = "abc", shifts = [[0,1,0],[1,2,1],[0,2,1]]
输出："ace"
解释：首先，将下标从 0 到 1 的字母向前移位，得到 s = "zac" 。
然后，将下标从 1 到 2 的字母向后移位，得到 s = "zbd" 。
最后，将下标从 0 到 2 的字符向后移位，得到 s = "ace" 。
示例 2:

输入：s = "dztz", shifts = [[0,0,0],[1,1,1]]
输出："catz"
解释：首先，将下标从 0 到 0 的字母向前移位，得到 s = "cztz" 。
最后，将下标从 1 到 1 的字符向后移位，得到 s = "catz" 。
*/
func shiftingLetters(s string, shifts [][]int) string {

	n := len(s)
	diff := make([]byte, n+1)
	diff[0] = s[0] - 'a'

	for i := 1; i < n; i++ {
		diff[i] = (s[i] + byte(26) - s[i-1]) % 26
	}

	for _, v := range shifts {
		start := v[0]
		end := v[1]
		dir := v[2]
		if dir == 0 {
			diff[start] = (byte(26) + diff[start] - byte(1)) % 26
			diff[end+1] = (diff[end+1] + byte(1)) % 26
		} else {
			diff[start] = (diff[start] + byte(1)) % 26
			diff[end+1] = (byte(26) + diff[end+1] - byte(1)) % 26
		}
	}
	ans := make([]byte, n)
	sum := byte(0)
	for i := 0; i < n; i++ {
		sum = (sum + diff[i]) % 26
		ans[i] = 'a' + sum
	}

	return string(ans)
}

func main() {
	println(shiftingLetters("dztz", [][]int{{0, 0, 0}, {1, 1, 1}}))
}
