package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var a uint32

	i := 31
	for i >= 0 && num != 0 {
		a = a ^ (num&1)<<i
		i--
		num >>= 1
	}
	return a
}
func main() {

	var a uint32
	fmt.Println(a << 31)
}
