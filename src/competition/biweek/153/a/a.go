package main

import "fmt"

func reverseDegree(s string) int {

	ans := 0

	for i := range s {

		ans += (26 - int(s[i]-'a')) * (i + 1)
	}
	return ans
}
func main() {
	fmt.Println(reverseDegree("abc"))
}
