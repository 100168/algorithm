package main

import (
	"fmt"
)

/*有效括号*/
func isValid(s string) bool {

	n := len(s)
	if n%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	var stack []byte

	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else if pairs[s[i]] == 0 {
			if pairs[stack[len(stack)-1]] != s[i] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
func main() {
	pairs := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	fmt.Println(pairs[1])

	fmt.Println(pairs[4])

	isValid("()")

}
