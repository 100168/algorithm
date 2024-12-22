package main

func numberOfSpecialChars(word string) int {
	cnt := make(map[uint8]bool)

	vis := make(map[uint8]bool)

	ans := 0
	for i := range word {
		if word[i] < word[i]^32 && cnt[word[i]^32] && !cnt[word[i]] {
			cnt[word[i]] = true
			vis[word[i]] = true
		}
		if word[i] > word[i]^32 {
			if cnt[word[i]^32] {
				vis[word[i]^32] = false
			}
		}
		cnt[word[i]] = true
	}

	for _, v := range vis {
		if v {
			ans++
		}
	}
	return ans
}

func main() {
	println(numberOfSpecialChars("DdD"))
}
