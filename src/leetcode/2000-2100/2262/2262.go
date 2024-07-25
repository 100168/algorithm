package main

import (
	"fmt"
	"math/bits"
)

/*
字符串的 引力 定义为：字符串中 不同 字符的数量。

例如，"abbca" 的引力为 3 ，因为其中有 3 个不同字符 'a'、'b' 和 'c' 。
给你一个字符串 s ，返回 其所有子字符串的总引力 。

子字符串 定义为：字符串中的一个连续字符序列。



示例 1：

输入：s = "abbca"
输出：28
解释："abbca" 的子字符串有：
- 长度为 1 的子字符串："a"、"b"、"b"、"c"、"a" 的引力分别为 1、1、1、1、1，总和为 5 。
- 长度为 2 的子字符串："ab"、"bb"、"bc"、"ca" 的引力分别为 2、1、2、2 ，总和为 7 。
- 长度为 3 的子字符串："abb"、"bbc"、"bca" 的引力分别为 2、2、3 ，总和为 7 。
- 长度为 4 的子字符串："abbc"、"bbca" 的引力分别为 3、3 ，总和为 6 。
- 长度为 5 的子字符串："abbca" 的引力为 3 ，总和为 3 。
引力总和为 5 + 7 + 7 + 6 + 3 = 28 。

*/

func appealSum(s string) int64 {

	type pair struct {
		v   int
		cnt int
	}
	st := make([]pair, 0)
	ans := 0
	for _, v := range s {
		c := int(v - 'a')
		st = append(st, pair{1 << c, 1})
		temp := st
		st = nil
		for _, p := range temp {
			p.v = p.v | 1<<c
			if len(st) > 0 && temp[len(st)-1].v == p.v {
				st[len(st)-1].cnt += p.cnt
			} else {
				st = append(st, p)
			}
			count := bits.OnesCount(uint(p.v))
			ans += count * p.cnt
		}

	}
	return int64(ans)
}

/*
*
太秀了
思路:贡献法
1.记录每个字符出现的最后一个下标
2.统计上次出现的下标到当前下标的距离
3.初始值为-1
*/
func appealSum2(s string) (ans int64) {
	sumG, last := 0, [26]int{}
	for i := range last {
		last[i] = -1
	} // 初始化成 -1 可以让提示 2-2 中的两种情况合并成一个公式
	for i, c := range s {
		c -= 'a'
		sumG += i - last[c]
		ans += int64(sumG)
		last[c] = i
	}
	return
}

func main() {
	fmt.Println(appealSum("abbca"))
	fmt.Println(appealSum2("abbca"))
}
