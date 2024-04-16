package main

import (
	"fmt"
)

func CF467B() {

}

func main() {
	var t, a, b int
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&a, &b)
		fmt.Println(a + b)
	}
}
