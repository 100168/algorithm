package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func CF1955B(_r io.Reader, out io.Writer) {

	var t, n, c, d int

	in := bufio.NewReader(_r)
	fmt.Fscan(in, &t)

next:
	for ; t > 0; t-- {
		_, _ = fmt.Fscan(in, &n, &c, &d)

		nums := make([]int, n*n)

		cntMap := make(map[int]int)

		for i := range nums {
			_, _ = fmt.Fscan(in, &nums[i])
			cntMap[nums[i]]++
		}
		sort.Ints(nums)

		memo := make([][]int, n)
		for i := range memo {
			memo[i] = make([]int, n)
		}
		memo[0][0] = nums[0]
		for i := 1; i < n; i++ {
			memo[i][0] = memo[i-1][0] + c
			cntMap[memo[i][0]]--
			if cntMap[memo[i][0]] < 0 {
				_, _ = fmt.Fprintln(out, "NO")
				continue next
			}
		}
		for j := 1; j < n; j++ {
			memo[0][j] = memo[0][j-1] + d
			cntMap[memo[0][j]]--
			if cntMap[memo[0][j]] < 0 {
				_, _ = fmt.Fprintln(out, "NO")
				continue next
			}
		}

		for i := 1; i < n; i++ {
			for j := 1; j < n; j++ {
				memo[i][j] = memo[i][j-1] + d
				cntMap[memo[i][j]]--
				if cntMap[memo[i][j]] < 0 {
					_, _ = fmt.Fprintln(out, "NO")
					continue next
				}
			}
		}
		_, _ = fmt.Fprintln(out, "YES")
	}

}

func main() {
	CF1955B(os.Stdin, os.Stdout)
}
