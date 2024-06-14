package main

import (
	"fmt"
	"sort"
)

/*
*
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

这是什么玩意
继续加油吧
*/
const mod = 100_000_000_7

func countKSubsequencesWithMaxBeauty(s string, k int) int {

	//先计数
	count := make([]int, 26)
	for _, v := range s {
		count[v-'a']++
	}

	type pair struct {
		cnt int
		num int
	}

	ans := 1
	cc := make(map[int]int)
	for _, v := range count {
		cc[v]++
	}
	kv := make([]pair, len(cc))

	for k, v := range cc {
		kv = append(kv, pair{cnt: k, num: v})
	}
	sort.Slice(kv, func(i, j int) bool {
		return kv[i].cnt > kv[j].cnt
	})
	for _, v := range kv {
		if v.num >= k {
			return ans * (pow(v.cnt, k) * comb(v.num, k) % mod) % mod
		}
		ans = (ans * pow(v.cnt, v.num)) % mod
		k -= v.num

	}
	//魅力值
	return 0

}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func comb(n, k int) int {
	res := n
	div := 1
	for i := 2; i <= k; i++ {
		res = res * (n - i + 1) % mod
		div = div * i % mod
	}

	return res * pow(div, mod-2) % mod
}

func main() {
	fmt.Println(countKSubsequencesWithMaxBeauty("xmfykrfjvnazcrytjpzmimmfqyutsmlzwiprlsyjvbadgpqogzttcrfbevjwnigmfiykajzjzuxbhumtfhlnvtbvuunvatzqknektuthhfhoiypqkciojaedonhgqagygkvtxvxbonncuqcynmxsompedarstrmxboqvqtnfksmlmswhlobesenebkmsijzrhonllglarndkjjdridqvmrmgjuikrtdaksmzdrrybggiezhsdkegtrhqzouwcfwkybtlewxndua", 18))
}
