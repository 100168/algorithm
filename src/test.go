package main

import "fmt"

func getBinary(num int) int {
	count := 0
	for num > 0 {
		count++
		num &= num - 1
	}
	return count
}
func main() {
	fmt.Println(1<<13 | 1<<12 | 1<<7 | 1<<6 | 1<<4 | 1<<2)
}
