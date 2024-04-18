package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func CF1352B(_r io.Reader, out io.Writer) {

	in := bufio.NewReader(_r)

	var t, n, k int
	fmt.Fscan(in, &t)
next:
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &k)

		ans := make([]int, 0)
		mod := n % k

		if mod == 0 {
			for j := 0; i < k; j++ {
				ans = append(ans, n/k)
			}
		} else {
			same := n / (k - 1)
			if same%2 == mod%2 {
				for i := 0; i < k-1; i++ {
					ans = append(ans, same)
				}
				ans = append(ans, mod)
			} else {
				fmt.Fprintln(out, "NO")
				continue next
			}

		}
		fmt.Fprintln(out, "YES")
		res := fmt.Sprint(ans)
		res = strings.Trim(res, "[")
		res = strings.Trim(res, "]")
		fmt.Fprintln(out, res)
	}
}

func main() {
	CF1352B(os.Stdin, os.Stdout)
}
