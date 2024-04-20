package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

func CF1537C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var t, n int

	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		fmt.Fscan(in, &n)
		h := make([]int, n)
		for i := range h {
			fmt.Fscan(in, &h[i])
		}
		sort.Ints(h)
		if n == 2 {
			fmt.Fprintln(out, h[0], h[1])
			continue
		}
		mn := math.MaxInt
		pos := 0
		for i := 1; i < n; i++ {
			if h[i]-h[i-1] < mn {
				mn = h[i] - h[i-1]
				pos = i
			}
		}
		for i := pos; i < n; i++ {
			fmt.Fprint(out, h[i], " ")
		}
		for i := 0; i < pos; i++ {
			fmt.Fprint(out, h[i], " ")
		}
		fmt.Fprintln(out)
	}
}

func main() {
	CF1537C(os.Stdin, os.Stdout)
}
