package main

import "fmt"

/*
*
给定整数 p 和 m ，一个长度为 k 且下标从 0 开始的字符串 s 的哈希值按照如下函数计算：

hash(s, p, m) = (val(s[0]) * p0 + val(s[1]) * p1 + ... + val(s[k-1]) * pk-1) mod m.
其中 val(s[i]) 表示 s[i] 在字母表中的下标，从 val('a') = 1 到 val('z') = 26 。

给你一个字符串 s 和整数 power，modulo，k 和 hashValue 。
请你返回 s 中 第一个 长度为 k 的 子串 sub ，满足 hash(sub, power, modulo) == hashValue 。

测试数据保证一定 存在 至少一个这样的子串。

子串 定义为一个字符串中连续非空字符组成的序列。

x*n*mod /p
*/

func subStrHash(s string, power, mod, k, hashValue int) (ans string) {
	n := len(s)
	// 用秦九韶算法计算 s[n-k:] 的哈希值，同时计算 pk=power^k
	hash, pk := 0, 1

	index := -1
	for i := n - 1; i >= n-k; i-- {
		hash = (hash*power + int(s[i]&31)) % mod
		pk = pk * power % mod
	}
	if hash == hashValue {
		index = n - k
	}
	for i := n - k - 1; i >= 0; i-- {

		//hash = (hash*power + int(s[i]&31) - pk*int(s[i+k]&31)%mod + mod) % mod
		hash = (hash*power + int(s[i]&31) - pk*int(s[i+k]&31)%mod + mod) % mod
		if hash == hashValue {
			index = i
		}
	}
	ans = s[index : index+k]
	return
}

func main() {
	fmt.Println(subStrHash("xxterzixjqrghqyeketqeynekvqhc", 15, 94, 4, 16))
	fmt.Println(subStrHash("leetcode", 7, 20, 2, 0))
	fmt.Println(subStrHash("dlojuxgftvpqpsknfgkejydsxgcgyroavsefjrejytcgflrnnxxsxowqbteycujnrbaokjibq", 8, 54, 30, 2))
}
