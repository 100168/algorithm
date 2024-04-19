package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF478B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)

	//5 1   4+3+2+1==10       5 2      1   4   3+2+1   2,3 1+2+1

	var n, m int

	fmt.Fscan(in, &n, &m)
	k := n / m
	mod := n % m
	minVal := 0
	minVal = (k+1)*(k)/2*(mod) + (k)*(k-1)/2*(m-mod)
	maxVal := (n - (m - 1)) * ((n - (m - 1)) - 1) / 2
	fmt.Fprintln(out, minVal, maxVal)
}
func main() {

	CF478B(os.Stdin, os.Stdout)
}
