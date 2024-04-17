package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
	"os"
)

func CF467B(_r io.Reader, out io.Writer) {

	in := bufio.NewReader(_r)
	var n, m, k, v, ans uint
	fmt.Fscan(in, &n, &m, &k)
	a := make([]uint, m)

	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	fmt.Fscan(in, &v)

	for _, w := range a {
		if uint(bits.OnesCount(v^w)) <= k {
			ans++
		}
	}
	fmt.Fprint(out, ans)
}

func main() {
	CF467B(os.Stdin, os.Stdout)
}
