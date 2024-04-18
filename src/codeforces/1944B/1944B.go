package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//1^2 == 2
//2^4 == 6

func CF1944B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)

	var t, n, k int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &k)
		nums := make([]int, 2*n)
		for j := range nums {
			fmt.Fscan(in, &nums[j])
		}
		left := nums[:n]
		cntLeft := make([]int, n+1)
		ansLeft := make([]int, 0)
		ansRight := make([]int, 0)
		for _, v := range left {
			cntLeft[v]++
		}

		for i, v := range cntLeft[1:] {
			if v == 0 && len(ansRight) < 2*k {
				ansRight = append(ansRight, i+1)
				ansRight = append(ansRight, i+1)
			}
			if v == 2 && len(ansLeft) < 2*k {
				ansLeft = append(ansLeft, i+1)
				ansLeft = append(ansLeft, i+1)
			}

		}
		for i, v := range cntLeft[1:] {
			if v == 1 && len(ansRight) < 2*k {
				ansRight = append(ansRight, i+1)
				ansLeft = append(ansLeft, i+1)
			}
		}
		l := fmt.Sprint(ansLeft)
		l = strings.Trim(l, "[")
		l = strings.Trim(l, "]")
		r := fmt.Sprint(ansRight)
		r = strings.Trim(r, "[")
		r = strings.Trim(r, "]")
		fmt.Fprintln(out, l)
		fmt.Fprintln(out, r)
	}
}

func main() {
	CF1944B(os.Stdin, os.Stdout)
}
