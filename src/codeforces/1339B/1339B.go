package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func CF1339B(_r io.Reader, out io.Writer) {

	in := bufio.NewReader(_r)

	var t, n int

	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		sort.Ints(a)

		m := (n - 1) / 2

		ans := make([]int, 0)

		// 5 -2 4 8 6 5
		//-2 4 5 5 6 8
		l, r := m, m+1

		for l >= 0 && r < n {
			ans = append(ans, a[l])
			ans = append(ans, a[r])
			l--
			r++
		}
		if l >= 0 {
			ans = append(ans, a[l])
			l--
		}
		res := fmt.Sprint(ans)
		res = strings.Trim(res, "[")
		res = strings.Trim(res, "]")

		fmt.Fprintln(out, res)
	}
}
func main() {

	CF1339B(os.Stdin, os.Stdout)

}
