package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//1 在集合中，x*a，x+b也在，
//  1  a = 3  b = 4   ,1,3,6,9,   5,9,3+4,7+4,
//  a*a   (a+b)*a+b = a**+(a+1)*b

func CF1542B(_r io.Reader, w io.Writer) {
	in := bufio.NewReader(_r)
	var t, n, a, b int

	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		fmt.Fscan(in, &n, &a, &b)

		if a == 1 {
			if (n-1)%b == 0 {
				fmt.Fprintln(w, "YES")
			} else {
				fmt.Fprintln(w, "NO")
			}
		} else {
			base := 1
			flag := 0
			for base <= n {
				if base%b == n%b {
					flag = 1
					break
				}
				base *= a
			}
			if flag == 1 {
				fmt.Fprintln(w, "YES")
			} else {
				fmt.Fprintln(w, "NO")
			}
		}
	}
}

func main() {
	CF1542B(os.Stdin, os.Stdout)
}
