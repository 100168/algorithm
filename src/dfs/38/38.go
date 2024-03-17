package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	ans := "1"
	for ; n > 0; n-- {
		ans = dfs(ans)
	}
	return ans

}

func dfs(s string) string {
	cur := s[0]

	n := len(s)
	ans := ""
	cnt := 1
	for i := 1; i < n; i++ {
		if s[i] != cur {
			ans += strconv.Itoa(cnt) + string(cur)
			cur = s[i]
			cnt = 1
		} else {
			cnt++
		}
	}
	return ans + strconv.Itoa(cnt) + string(cur)

}
func main() {
	fmt.Println(countAndSay(4))
	fmt.Println(-1 % 10)
}
