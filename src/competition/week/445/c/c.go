package main

import (
	"math/big"
	"sort"
)

func smallestPalindrome(s string, k int) string {
	cnt := make(map[rune]int)
	for _, c := range s {
		cnt[c]++
	}
	n := len(s)
	var midChar rune

	if n%2 != 0 {
		var oddChar rune
		for c, v := range cnt {
			if v%2 == 1 {
				oddChar = c
				break
			}
		}
		midChar = oddChar
		cnt[midChar]--
		if cnt[midChar] == 0 {
			delete(cnt, midChar)
		}
		for c := range cnt {
			cnt[c] = cnt[c] / 2
			if cnt[c] == 0 {
				delete(cnt, c)
			}
		}
	} else {
		for c := range cnt {
			cnt[c] = cnt[c] / 2
			if cnt[c] == 0 {
				delete(cnt, c)
			}
		}
	}

	var chars []rune
	var keys []rune
	for c := range cnt {
		keys = append(keys, c)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, c := range keys {
		for i := 0; i < cnt[c]; i++ {
			chars = append(chars, c)
		}
	}
	m := len(chars)

	if m == 0 {
		possible := 1
		if k > possible {
			return ""
		} else {
			if midChar != 0 {
				return string(midChar)
			} else {
				return ""
			}
		}
	}

	total := big.NewInt(1)
	for i := 2; i <= m; i++ {
		total.Mul(total, big.NewInt(int64(i)))
	}
	freq := make(map[rune]int)
	for _, c := range chars {
		freq[c]++
	}
	for _, v := range freq {
		fact := big.NewInt(1)
		for i := 2; i <= v; i++ {
			fact.Mul(fact, big.NewInt(int64(i)))
		}
		total.Div(total, fact)
	}

	if k > int(total.Int64()) {
		return ""
	}

	var uniqueChars []rune
	for c := range freq {
		uniqueChars = append(uniqueChars, c)
	}
	sort.Slice(uniqueChars, func(i, j int) bool {
		return uniqueChars[i] < uniqueChars[j]
	})
	currentFreq := make(map[rune]int)
	for k, v := range freq {
		currentFreq[k] = v
	}
	var res []rune
	remaining := m
	currentTotal := new(big.Int).Set(total)

	for i := 0; i < m; i++ {
		for _, c := range uniqueChars {
			if currentFreq[c] == 0 {
				continue
			}
			num := new(big.Int).Set(currentTotal)
			num.Mul(num, big.NewInt(int64(currentFreq[c])))
			num.Div(num, big.NewInt(int64(remaining)))
			if k > int(num.Int64()) {
				k -= int(num.Int64())
				continue
			} else {
				res = append(res, c)
				currentTotal = num
				currentFreq[c]--
				remaining--
				break
			}
		}
	}

	prefix := string(res)
	if midChar != 0 {
		return prefix + string(midChar) + rs(prefix)
	} else {
		return prefix + rs(prefix)
	}
}

func rs(s string) string {
	bs := []byte(s)
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}
