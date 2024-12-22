package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func do(_r io.Reader, out io.Writer) {

	in := bufio.NewScanner(_r)

	in.Scan()
	in.Text()
	in.Scan()
	s1 := in.Text()
	in.Scan()
	s2 := in.Text()
	cnt1 := make([]int, 26)
	cnt2 := make([]int, 26)

	for i := range s1 {
		cnt1[s1[i]-'a']++
	}
	for i := range s2 {
		if cnt1[s2[i]-'a'] > 0 {
			cnt1[s2[i]-'a']--
		} else {
			cnt2[s2[i]-'a']++
		}
	}

	ans := 0
	p1 := 0
	for i := 0; i < 26; i++ {

		for cnt1[i] > 0 {
			for cnt2[p1] == 0 {
				p1++
			}
			cnt2[p1]--
			ans += abs(i - p1)
			cnt1[i]--
		}

	}

	fmt.Fprintln(out, ans)

}

func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

func main() {
	do(os.Stdin, os.Stdout)
}
