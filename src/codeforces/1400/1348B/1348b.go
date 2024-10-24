package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
)

func CF1348b(_r io.Reader, out io.Writer) {

	in := bufio.NewReader(_r)

	var t, n, k int

	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &nums[i])
		}

		newNums := slices.Clone(nums)

		sort.Ints(newNums)
		newNums = slices.Compact(newNums)

		ans := make([]int, 0)

		//重复元素大于k直接返回-1
		if len(newNums) > k {
			fmt.Fprintln(out, -1)
		} else {
			//否则先填充k个
			//然后再打印n次
			for len(newNums) < k {
				newNums = append(newNums, 1)
			}
			fmt.Fprintln(out, n*k)
			for ; n > 0; n-- {
				ans = append(ans, slices.Clone(newNums)...)
			}

			for _, v := range ans {
				fmt.Fprint(out, v, " ")
			}

		}

	}
	// 1 2
}

func main() {

	//
	CF1348b(os.Stdin, os.Stdout)
}
