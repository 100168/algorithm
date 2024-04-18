package main

import "math"

/*
小扣出去秋游，途中收集了一些红叶和黄叶，他利用这些叶子初步整理了一份秋叶收藏集 leaves， 字符串 leaves 仅包含小写字符 r 和 y，
其中字符 r 表示一片红叶，字符 y 表示一片黄叶。出于美观整齐的考虑，小扣想要将收藏集中树叶的排列调整成「红、黄、红」三部分。
每部分树叶数量可以不相等，但均需大于等于 1。每次调整操作，小扣可以将一片红叶替换成黄叶或者将一片黄叶替换成红叶。
请问小扣最少需要多少次调整操作才能将秋叶收藏集调整完毕。
*/
func minimumOperations(leaves string) int {

	n := len(leaves)

	a, b, c := 0, math.MaxInt/2, math.MaxInt/2

	for i := 0; i < n; i++ {
		cur := leaves[i]
		if cur == 'r' {
			t := b
			if i > 0 {
				b = min(a, b) + 1

			}
			if i > 1 {
				c = min(t, c)
			}
		} else {
			t1 := a
			t2 := b
			a += 1
			if i > 0 {
				b = min(t1, b)
			}
			if i > 1 {
				c = min(t2, c) + 1
			}
		}
	}
	return c
}
