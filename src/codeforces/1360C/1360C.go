package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF1360C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)

	var t, n int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		numsMap := make(map[int]bool)

		flag := false
		cntOdd := 0
		for i := 0; i < n; i++ {
			if a[i]%2 != 0 {
				cntOdd++
			}
			numsMap[a[i]] = true
			if numsMap[a[i]-1] || numsMap[a[i]+1] {
				flag = true
			}
		}
		if cntOdd%2 == 0 {
			fmt.Fprintln(out, "YES")
		} else if flag {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func main() {
	CF1360C(os.Stdin, os.Stdout)
}
