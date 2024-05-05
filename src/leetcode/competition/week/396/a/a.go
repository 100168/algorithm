package main

func isValid(word string) bool {

	cntNumb := 0
	cntY := 0
	cntOther := 0

	nums := map[uint8]bool{'1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true}
	//'a'、'e'、'i'、'o'、'u' 及其大写形式都属于 元音字母 。
	chars := map[uint8]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	for i := range word {

		cur := word[i]
		if nums[cur] {
			cntNumb++
		} else if chars[cur] || chars[cur^32] {
			cntY++
		} else if cur != '@' && cur != '#' && cur != '$' {
			cntOther++
		} else {
			return false
		}

	}

	return cntY > 0 && cntOther > 0 && len(word) >= 3
}
