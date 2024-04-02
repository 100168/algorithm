package main

import (
	"sort"
)

/*
给你一个字符串 s 和一个整数 k 。

k 子序列指的是 s 的一个长度为 k 的 子序列 ，且所有字符都是 唯一 的，也就是说每个字符在子序列里只出现过一次。

定义 f(c) 为字符 c 在 s 中出现的次数。

k 子序列的 美丽值 定义为这个子序列中每一个字符 c 的 f(c) 之 和 。

比方说，s = "abbbdd" 和 k = 2 ，我们有：

f('a') = 1, f('b') = 3, f('d') = 2
s 的部分 k 子序列为：
"abbbdd" -> "ab" ，美丽值为 f('a') + f('b') = 4
"abbbdd" -> "ad" ，美丽值为 f('a') + f('d') = 3
"abbbdd" -> "bd" ，美丽值为 f('b') + f('d') = 5
请你返回一个整数，表示所有 k 子序列 里面 美丽值 是 最大值 的子序列数目。由于答案可能很大，将结果对 109 + 7 取余后返回。

一个字符串的子序列指的是从原字符串里面删除一些字符（也可能一个字符也不删除），不改变剩下字符顺序连接得到的新字符串。

注意：

f(c) 指的是字符 c 在字符串 s 的出现次数，不是在 k 子序列里的出现次数。
两个 k 子序列如果有任何一个字符在原字符串中的下标不同，则它们是两个不同的子序列。所以两个不同的 k 子序列可能产生相同的字符串。
*/

const mod = 100_000_000_7

func countKSubsequencesWithMaxBeauty(s string, k int) int {

	count := make([]int, 26)
	for _, v := range s {
		count[v-'a']++
	}

	type KV struct {
		cnt int
		num int
	}

	ans := 1
	cc := make(map[int]int, 26)
	for _, i2 := range count {
		cc[i2]++
	}
	kv := make([]KV, len(cc))

	for k, v := range cc {
		kv = append(kv, KV{cnt: k, num: v})
	}
	sort.Slice(kv, func(i, j int) bool {
		return kv[i].cnt > kv[j].cnt
	})
	for _, v := range kv {
		if v.num >= k {
			return ans * (pow(v.cnt, k) % mod * comb(v.num, k)) % mod
		}
		ans = (ans * pow(v.cnt, v.num)) % mod
		k -= v.num

	}
	//魅力值
	return 0

}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func comb(n, k int) int {
	res := n
	for i := 2; i <= k; i++ {
		res = res * (n - i + 1) / i // n,n-1,n-2,... 中的前 i 个数至少有一个因子 i
	}
	return res % mod
}

func main() {
	//println(countKSubsequencesWithMaxBeauty("bcca", 2))
	println(countKSubsequencesWithMaxBeauty("abbcd", 4))
}
