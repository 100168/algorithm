package main

import "fmt"

func calculateScore(s string) int64 {

	st := make([][]int, 26)

	ans := 0

	for i := range s {

		c := int(s[i] - 'a')

		rs := 25 - c

		if len(st[rs]) != 0 {
			pre := st[rs][len(st[rs])-1]
			st[rs] = st[rs][:len(st[rs])-1]
			ans += i - pre
		}
		st[c] = append(st[c], i)
	}

	return int64(ans)
}
func main() {

	fmt.Println(calculateScore("eockppxdqclkhjgvnw"))
}
