package main

func numberOfSpecialChars(word string) int {

	cnt := make(map[uint8]bool)

	ans := 0
	for i := range word {
		if cnt[word[i]^32] && !cnt[word[i]] {
			ans++
		}
		cnt[word[i]] = true
	}

	return ans
}
