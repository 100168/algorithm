package main

import "fmt"

func alienOrder(words []string) string {

	in := make([]int, 26)

	g := make([]map[int]bool, 26)

	exist := make(map[int]bool)

	for i := range g {
		g[i] = make(map[int]bool)
	}

	for _, v := range words {

		for _, x := range v {

			y := int(byte(x) - 'a')
			exist[y] = true

		}
	}
	for j, v := range words[1:] {

		pre := words[j]

		for i := 0; i < min(len(pre), len(v)); i++ {
			x, y := int(pre[i]-'a'), int(v[i]-'a')
			if x != y {
				if !g[x][y] {
					in[y]++
					g[x][y] = true
				}
				break
			}

		}

		if len(pre) > len(v) && pre[:len(v)] == v {
			return ""
		}

	}

	var st []int

	var ans []byte
	for i, v := range in {

		if v == 0 && exist[i] {
			st = append(st, i)

		}
	}

	for len(st) > 0 {
		t := st[0]
		st = st[1:]
		ans = append(ans, byte(t+'a'))

		for v, _ := range g[t] {
			in[v]--
			if in[v] == 0 {
				st = append(st, v)

			}
		}

	}

	for _, v := range in {

		if v > 0 {
			return ""
		}
	}
	return string(ans)

}
func main() {
	fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt", "te"}))
}
