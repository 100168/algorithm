package main

import "fmt"

/*
*
Alice 和 Bob 正在玩一个幻想战斗游戏，游戏共有 n 回合，
每回合双方各自都会召唤一个魔法生物：火龙（F）、水蛇（W）或地精（E）。每回合中，
双方 同时 召唤魔法生物，并根据以下规则得分：

如果一方召唤火龙而另一方召唤地精>召唤 火龙 的玩家将获得一分。
如果一方召唤水蛇而另一方召唤火龙>召唤 水蛇 的玩家将获得一分。
如果一方召唤地精而另一方召唤水蛇>召唤 地精 的玩家将获得一分。
如果双方召唤相同的生物，那么两个玩家都不会获得分数。
给你一个字符串 s，包含 n 个字符 'F'、'W' 和 'E'，代表 Alice 每回合召唤的生物序列：

如果 s[i] == 'F'，Alice 召唤火龙。
如果 s[i] == 'W'，Alice 召唤水蛇。
如果 s[i] == 'E'，Alice 召唤地精。
Bob 的出招序列未知，但保证 Bob 不会在连续两个回合中召唤相同的生物。如果在 n 轮后 Bob 获得的总分 严格大于 Alice 的总分，则 Bob 战胜 Alice。

返回 Bob 可以用来战胜 Alice 的不同出招序列的数量。

由于答案可能非常大，请返回答案对 109 + 7 取余 后的结果。

示例 1：

输入： s = "FFF"

输出： 3

解释：

Bob 可以通过以下 3 种出招序列战胜 Alice："WFW"、"FWF" 或 "WEW"。注意，其他如 "WWE" 或 "EWW" 的出招序列是无效的，因为 Bob 不能在连续两个回合中使用相同的生物。

示例 2：

输入： s = "FWEFW"

输出： 18

解释：

Bob 可以通过以下出招序列战胜 Alice："FWFWF"、"FWFWE"、"FWEFE"、"FWEWE"、"FEFWF"、"FEFWE"、"FEFEW"、"FEWFE"、"WFEFE"、"WFEWE"、"WEFWF"、"WEFWE"、"WEFEF"、"WEFEW"、"WEWFW"、"WEWFE"、"EWFWE" 或 "EWEWE"。

提示：

1 <= s.length <= 1000
s[i] 是 'F'、'W' 或 'E' 中的一个。
*/
func countWinningSequences(s string) int {

	mod := int(1e9 + 7)

	n := len(s)
	f := make([][]map[byte]int, n)

	for i := range f {
		f[i] = make([]map[byte]int, 2*n+1)
		for j := range f[i] {
			f[i][j] = make(map[byte]int)
		}
	}

	type pair struct {
		bob   byte
		alice byte
	}

	//F>E
	//W>F
	//E>W

	pairs := []pair{{'F', 'E'}, {'W', 'F'}, {'E', 'W'}}

	pairsMap := make(map[pair]int)

	for _, v := range pairs {
		pairsMap[v] = 1
	}

	chars := []byte{'E', 'F', 'W'}

	var dfs func(int, int, byte) int

	//ss := ""

	dfs = func(i int, cnt int, pre byte) int {

		if i < 0 {
			if cnt > n {
				return 1
			}
			return 0
		}

		if cnt+i+1 < n {
			return 0
		}

		if _, ok := f[i][cnt][pre]; ok {
			return f[i][cnt][pre]
		}
		alice := s[i]

		cur := 0
		for _, v := range chars {
			if v == pre {
				continue
			}
			//ss = string(v) + ss
			if v == alice {
				cur += dfs(i-1, cnt, v)
				//ss = ss[1:]
				continue
			}
			if pairsMap[pair{v, alice}] == 1 {
				cur += dfs(i-1, cnt+1, v)
				//ss = ss[1:]
				continue
			}
			cur += dfs(i-1, cnt-1, v)
			//ss = ss[1:]
		}
		cur %= mod
		f[i][cnt][pre] = cur
		return cur
	}
	return dfs(n-1, n, 1)
}

func main() {

	//F>E
	//W>F
	//E>W

	fmt.Println(countWinningSequences("FFW"))
	//EFE

}
