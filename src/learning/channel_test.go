package learning

import (
	"fmt"
	"testing"
	"time"
)

func sayHello(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

func TestConcurrent(t *testing.T) {
	go sayHello("hello")
	sayHello("world")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func TestChannel(t *testing.T) {

	c := make(chan int)

	s := []int{1, 2, 3, 4, 5}
	go sum(s[:2], c)
	go sum(s[1:4], c)
	x := <-c

	b := make(chan int, 2)
	b <- 2
	b <- 1
	fmt.Println(<-b)
	fmt.Println(<-b)
	fmt.Println(x)
}

func producer(s chan int) {
	for i := 0; i < 50; i++ {
		time.Sleep(time.Second)
		s <- i
	}

}

func TestSelect(t *testing.T) {

	s1 := make(chan int, 10)
	s2 := make(chan int, 10)
	go producer(s1)
	go producer(s2)

	for {
		select {
		case a := <-s1:
			fmt.Println("s1", a)
		case b := <-s2:
			fmt.Println("s2", b)
		}
	}

}
