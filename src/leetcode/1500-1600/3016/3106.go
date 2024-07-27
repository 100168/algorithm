package main

import "fmt"

func getSmallestString(s string, k int) string {

	bytes := []byte(s)

	for i := range s {

		cur := int(s[i] - 'a')
		diff := min(cur-0, 26-cur)
		if k >= diff {
			bytes[i] = 'a'
		} else {
			bytes[i] = byte(min(cur-k, (cur+k)%26)) + 'a'
		}
		k -= diff
		if k <= 0 {
			break
		}
	}

	return string(bytes)
}

func main() {
	fmt.Println(getSmallestString("uo", 5))
}
