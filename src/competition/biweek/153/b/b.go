package main

import "fmt"

/*
*给你一个长度为 n 的二进制字符串 s，其中：

'1' 表示一个 活跃 区段。
'0' 表示一个 非活跃 区段。
你可以执行 最多一次操作来最大化 s 中的活跃区段数量。在一次操作中，你可以：

将一个被 '0' 包围的连续 '1' 区块转换为全 '0'。
然后，将一个被 '1' 包围的连续 '0' 区块转换为全 '1'。
返回在执行最优操作后，s 中的 最大 活跃区段数。

注意：处理时需要在 s 的两侧加上 '1' ，即 t = '1' + s + '1'。这些加上的 '1' 不会影响最终的计数。

示例 1：

输入： s = "01"

输出： 1

解释：

因为没有被 '0' 包围的 '1' 区块，因此无法进行有效操作。最大活跃区段数为 1。©leetcode
*/
func maxActiveSectionsAfterTrade(s string) int {

	s = "1" + s + "1"
	n := len(s)

	pre := make([]int, n)
	suf := make([]int, n)

	ans := 0
	pre[0] = 0
	suf[n-1] = 0
	for i := 1; i < n-1; i++ {
		if s[i] == '1' {
			if s[i-1] == s[i] {
				pre[i] = pre[i-1] + 1
			} else {
				pre[i] = 1
			}
			ans++
		}
	}
	for i := n - 2; i > 0; i-- {
		if s[i] == '1' {
			if s[i+1] == s[i] {
				suf[i] = suf[i+1] + 1
			} else {
				suf[i] = 1
			}
		}

	}

	cnt := 0
	for i := 1; i < n-1; i++ {
		if s[i] == '1' {
			continue
		}
		j := i + 1

		for j < n && s[j] == s[i] && s[j] == '0' {
			j++
		}

		cc := j - i

		for j < n-1 && s[j] == '1' {
			j++
		}

		t := j

		for j < n && s[j] == '0' {
			j++
		}

		cc += j - t

		if j == t {
			break
		}

		cnt = max(cnt, cc)

		i = t - 1

	}
	return ans + cnt

}

func main() {

	//fmt.Println(maxActiveSectionsAfterTrade("01"))
	//fmt.Println(maxActiveSectionsAfterTrade("0100"))
	fmt.Println(maxActiveSectionsAfterTrade("000000"))
}
