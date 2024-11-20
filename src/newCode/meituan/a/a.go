package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func a(_r io.Reader, out io.Writer) {

	in := bufio.NewReader(_r)
	var n int
	fmt.Fscan(in, &n)

	nums := make([][]int, n)

	for i := range nums {
		nums[i] = make([]int, n)
	}

	sum := make([][]int, n+1)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		var c string
		fmt.Fscan(in, &c)

		for j := 0; j < n; j++ {
			nums[i][j] = int(c[j] - '0')
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] + nums[i][j] - sum[i][j]
			for k := 2; k <= n; k++ {
				if i+1 >= k && j+1 >= k {
					s := sum[i+1][j+1] - sum[i+1-k][j+1] - sum[i+1][j+1-k] + sum[i+1-k][j+1-k]
					if s*2 == k*k {
						ans[k-1]++
					}
				}

			}
		}
	}

	for _, v := range ans {
		fmt.Fprintln(out, v)
	}

}

func main() {
	a(os.Stdin, os.Stdout)
}
