package main

func numberOfChild(n int, k int) int {

	mul := k / (n - 1)
	mod := k % (n - 1)
	if mul%2 == 0 {
		return mod
	}
	return n - 1 - mod

}
