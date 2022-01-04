package main

import "fmt"

func decodeString(s string) string {
	var stack []byte
	for i := range s {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			peek := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var temp []byte
			for peek != '[' {
				temp = append(temp, peek)
				if len(stack) > 0 {
					peek, stack = stack[len(stack)-1], stack[:len(stack)-1]
				} else {
					break
				}
			}
			peek, stack = stack[len(stack)-1], stack[:len(stack)-1]
			num := 0
			flag := 1
			for peek >= '0' && peek <= '9' {
				num += flag * int(peek-'0')
				flag *= 10
				if len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
					peek, stack = stack[len(stack)-1], stack[:len(stack)-1]
				} else {
					break
				}
			}
			for i := 0; i < num; i++ {
				for j := len(temp) - 1; j >= 0; j-- {
					stack = append(stack, temp[j])
				}
			}
		}
	}
	return string(stack)
}

func main() {
	a := "3[a]2[bc]"

	s := decodeString(a)

	fmt.Println(s)
}
