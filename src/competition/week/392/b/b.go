package main

import "fmt"

func getSmallestString(s string, k int) string {

	n := len(s)
	ans := make([]int, n)

	res := make([]uint8, n)
	for i := 0; i < n; i++ {
		ans[i] = int(s[i] - 'a')
	}

	for i := 0; i < n; i++ {
		diff := min(ans[i], 26-ans[i], k)

		if ans[i]+diff != 26 {
			ans[i] = (ans[i] - diff) % 26
		} else {
			ans[i] = (ans[i] + diff) % 26
		}
		k -= diff
		res[i] = uint8(ans[i]) + 'a'
	}
	return string(res)
}

func main() {

	fmt.Println(getSmallestString("xaxcd", 4))
}
