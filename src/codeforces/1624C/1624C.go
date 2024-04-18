package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func CF1624C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)

	var t, n int
	fmt.Fscan(in, &t)

next:
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[j])
		}
		sort.Ints(a)

		vis := make([]bool, n+1)
		//1,2,8,25
		for _, v := range a {
			for v > n || vis[v] == true {
				v /= 2
			}
			if v == 0 {
				fmt.Fprintln(out, "NO")
				continue next
			}
			vis[v] = true
		}
		fmt.Fprintln(out, "YES")
	}
}

func main() {
	CF1624C(os.Stdin, os.Stdout)
}
