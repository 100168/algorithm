package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func CF1942B(_r io.Reader, out io.Writer) {

	var t, n int

	in := bufio.NewReader(_r)

	fmt.Fscan(in, &t)

	// p   0 1 4 2 3
	// a = 1 1 -2 1 2
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		ans := make([]int, n)
		vis := make([]bool, n+1)
		mex := 0
		for i := range a {
			fmt.Fscan(in, &a[i])
			if a[i] > 0 {
				ans[i] = mex
			} else {
				ans[i] = mex - a[i]
			}
			vis[ans[i]] = true
			for vis[mex] {
				mex++
			}
		}
		res := fmt.Sprint(ans)
		res = strings.Trim(res, "[")
		res = strings.Trim(res, "]")
		fmt.Fprintln(out, res)
	}

}

func main() {
	CF1942B(os.Stdin, os.Stdout)
}
