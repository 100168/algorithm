package main

/*给你一个数组 arr ，数组中有 n 个 非空 字符串。

请你求出一个长度为 n 的字符串 answer ，满足：

answer[i] 是 arr[i] 最短 的子字符串，且它不是 arr 中其他任何字符串的子字符串。如果有多个这样的子字符串存在，
answer[i] 应该是它们中字典序最小的一个。如果不存在这样的子字符串，answer[i] 为空字符串。
请你返回数组 answer 。*/

func shortestSubstrings(arr []string) []string {
	n := len(arr)
	ans := make([]string, n)

	//用来表示排名第几的是哪个字符
	sa := make([][]int, n)
	//用来表示字符排名第几
	rk := make([][]int, n)
	for i := range sa {
		sa[i] = make([]int, 100000)
	}
	for i := range rk {
		rk[i] = make([]int, 100000)
	}

	for x := range arr {
		s := arr[x]
		m := len(s)
		//第一次计算
		cnt := make([]int, 26)
		for i := 0; i < m; i++ {
			cnt[s[i]-'a']++
		}
		for i := 1; i < 26; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			cur := s[i] - 'a'
			//排名cnt[cur]-1 为cur
			sa[x][cnt[cur]-1] = i
			//第i个字符排名为cnt[cur]-1
			rk[x][i] = cnt[cur] - 1
			cnt[cur]--
		}
		//倍增计算
		for exp := 1; 1<<exp < m; exp <<= 1 {
			y := make([]int, len(sa))
			z := make([]int, len(rk))
			copy(sa[x], y)
			copy(rk[x], z)
			cnt := make([]int, 26)
			for i := 0; i < n; i++ {
				cur := s[i+exp] - 'a'
				cnt[cur]++
			}
			for i := 1; i < 26; i++ {
				cnt[i] += cnt[i-1]
			}
			for i := n - 1; i >= 0; i-- {
				cur := s[i+exp] - 'a'
				y[rk[x][cnt[cur]-1]] = i
				cnt[cur]--
			}
			copy(y, sa[x])
		}
	}
	return ans
}

func main() {
	println(shortestSubstrings([]string{"cab", "ad", "bad", "c"}))
}
