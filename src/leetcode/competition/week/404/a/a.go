package main

import "fmt"

func maxHeightOfTriangle(red int, blue int) int {

	minH := func(a, b int) int {

		h := 0
		for i := 1; a >= i; i += 2 {
			h++
			if b < i+1 {
				return h
			}
			h++
			b -= i + 1
			a -= i
		}
		return h
	}
	return max(minH(red, blue), minH(blue, red))

}

func maxHeightOfTriangle2(red int, blue int) int {

	minH := func(a, b int) int {

		left := []int{b, a}
		i := 1
		for left[i%2] >= i {
			left[i%2] -= i
			i++
		}
		return i - 1
	}
	return max(minH(red, blue), minH(blue, red))

}
func main() {
	fmt.Println(maxHeightOfTriangle(2, 3))
	fmt.Println(maxHeightOfTriangle2(2, 3))
}
