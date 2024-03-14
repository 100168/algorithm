package main

import "strings"

/*给你一个数组 arr ，数组中有 n 个 非空 字符串。

请你求出一个长度为 n 的字符串 answer ，满足：

answer[i] 是 arr[i] 最短 的子字符串，且它不是 arr 中其他任何字符串的子字符串。如果有多个这样的子字符串存在，
answer[i] 应该是它们中字典序最小的一个。如果不存在这样的子字符串，answer[i] 为空字符串。
请你返回数组 answer 。*/

func shortestSubstrings(arr []string) []string {
	n := len(arr)
	ans := make([]string, n)
	for i, v := range arr {
		minV := ""
		for j := 1; j <= len(v) && len(minV) == 0; j++ {
			for z := 0; z+j <= len(v); z++ {
				isV := true
				cur := v[z : z+j]
				for y, x := range arr {
					if y == i {
						continue
					}
					if strings.Contains(x, cur) {
						isV = false
						break
					}
				}
				if isV {
					if minV == "" {
						minV = cur
					} else {
						if strings.Compare(minV, cur) > 0 {
							minV = cur
						}
					}
				}
			}
		}
		ans[i] = minV
	}
	return ans
}

func main() {
	println(shortestSubstrings([]string{"cab", "ad", "bad", "c"}))
}
