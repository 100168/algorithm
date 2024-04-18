package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF1352B(_r io.Reader, out io.Writer) {

	in := bufio.NewReader(_r)

	var t, n, k int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &k)

		n1 := n - (k - 1)
		if n1 > 0 && n1%2 == 1 {
			fmt.Fprintln(out, "YES")
			for j := 0; j < k-1; j++ {
				fmt.Fprint(out, 1, " ")
			}
			fmt.Fprintln(out, n1)
			continue
		}
		n2 := n - 2*(k-1)
		if n2 > 0 && n2%2 == 0 {
			fmt.Fprintln(out, "YES")
			for j := 0; j < k-1; j++ {
				fmt.Fprint(out, 2, " ")
			}
			fmt.Fprintln(out, n2)

			continue
		}
		fmt.Fprintln(out, "NO")
	}
}

func main() {
	CF1352B(os.Stdin, os.Stdout)
}
