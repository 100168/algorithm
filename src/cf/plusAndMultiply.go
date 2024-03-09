package main

func plusAndMultiply(n int, a int, b int) string {

	if n%a == 0 || (n-b)%a == 0 || n%b == 1 {
		return "Yes"
	}

	return "No"
}

func main() {

	n := 19260817
	a := 233
	b := 264
	println(plusAndMultiply(n, a, b))
}
