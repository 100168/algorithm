package main

import (
	"strings"
)

func spellchecker(wordlist []string, queries []string) []string {

	//('a', 'e', 'i', 'o', 'u')
	autoMap := map[int32]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	wordMap := make(map[string]bool)
	lowerCaseMap := make(map[string][]string)
	changeMap := make(map[string][]string)
	for _, v := range wordlist {
		wordMap[v] = true
		lower := strings.ToLower(v)
		lowerCaseMap[lower] = append(lowerCaseMap[lower], v)
		quote := []int32(lower)
		for j, cur := range quote {
			if autoMap[cur] {
				quote[j] = '*'

			}
		}
		newStr := string(quote)
		changeMap[newStr] = append(changeMap[newStr], v)
	}
	ans := make([]string, 0)
	for _, v := range queries {
		if wordMap[v] {
			ans = append(ans, v)
		} else if e, ok := lowerCaseMap[strings.ToLower(v)]; ok {
			ans = append(ans, e[0])
		} else {
			quote := []int32(strings.ToLower(v))
			for j, cur := range quote {
				if autoMap[cur] {
					quote[j] = '*'
				}
			}
			newStr := string(quote)
			if t, ok := changeMap[newStr]; ok {
				ans = append(ans, t[0])
			} else {
				ans = append(ans, "")
			}
		}
	}
	return ans
}

func main() {
	println(spellchecker([]string{"KiTe", "kite", "hare", "Hare"}, []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}))
}
