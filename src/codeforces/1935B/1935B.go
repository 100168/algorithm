package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF1935B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var t, n int

	_, _ = fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		_, _ = fmt.Fscan(in, &n)
		nums := make([]int, n)
		for i := range nums {
			fmt.Fscan(in, &nums[i])
		}
	}

}

func main() {
	CF1935B(os.Stdin, os.Stdout)
}
