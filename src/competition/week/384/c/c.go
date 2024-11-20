package main

import (
	"sort"
)

func maxPalindromesAfterOperations(words []string) int {

	cntMap := make(map[uint8]int)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			cntMap[word[i]]++
		}
	}
	cntRank := make([]int, 0)

	for _, cnt := range cntMap {
		cntRank = append(cntRank, cnt)
	}

	sort.Ints(cntRank)
	ans := 0

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for _, w := range words {
		n := len(w)

		cnt := 0
		for _, v := range cntRank {
			cnt += v / 2
		}

		if n%2 == 0 {
			if cnt >= n/2 {
				cnt = n / 2
				for i, v := range cntRank {
					cntRank[i] -= min(v/2*2, cnt*2)
					cnt -= max(0, v/2)
					if cnt == 0 {
						break
					}
				}
				ans++
			}
		} else {

			if cnt >= n/2 {
				if cnt >= n/2 && len(cntRank) >= n/2+1 {

					cnt = n / 2
					for i, v := range cntRank {
						cntRank[i] -= min(v/2*2, cnt*2)
						cnt -= max(0, v/2)
						if cnt == 0 {
							break
						}
					}
				}
				sort.Ints(cntRank)

				flag := false
				for i, v := range cntRank {
					if v%2 == 1 {
						cntRank[i]--
						flag = true
						break
					}
				}
				if !flag {
					for i, v := range cntRank {
						if v != 0 {
							cntRank[i]--
							break
						}
					}
				}
				ans++
			}

		}

	}
	return ans

}

func main() {
	println(maxPalindromesAfterOperations2([]string{"fchf", "fcca", "dcgb", "fgdg", "dabh"}))
}

func maxPalindromesAfterOperations2(words []string) int {

	cntMap := make(map[uint8]int)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			cntMap[word[i]]++
		}
	}
	odd := 0
	even := 0
	for _, cnt := range cntMap {
		if cnt%2 == 0 {
			even += cnt
		} else {
			odd++
			even += cnt - 1
		}
	}
	ans := 0

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for _, w := range words {
		n := len(w)

		if n%2 == 0 {
			if even >= n {
				ans++
				even -= n
			}
		} else {
			odd -= 1
			if odd < 0 {
				even -= 2
				odd += 2
			}
			if even >= n-1 {
				even -= n - 1
				ans++
			}

		}
	}
	return ans

}
