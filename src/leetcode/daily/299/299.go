package main

import (
	"fmt"
)

func getHint(secret string, guess string) string {

	cnt := make(map[uint8]int)

	for i := 0; i < len(secret); i++ {
		cnt[secret[i]]++
	}

	cntA := 0
	cntB := 0

	for i := 0; i < len(guess); i++ {
		if guess[i] == secret[i] {
			cntA++
			if cnt[guess[i]] < 0 {
				cntB--
			}
			cnt[guess[i]]--
		} else {
			if cnt[guess[i]] > 0 {
				cntB++
				cnt[guess[i]]--
			}
		}
	}
	return fmt.Sprintf("%dA%dB", cntA, cntB)
}

func main() {
	println(getHint("12342", "23132"))
}
