package main

import "fmt"

func main() {

	y, x := 100, 110
	fmt.Println(x + y)
	fmt.Println("hello world!")
	test(0)
}

func test(a int) {
	if a > 10 {
		return
	}
	fmt.Println("你好")
	test(a + 1)

}
