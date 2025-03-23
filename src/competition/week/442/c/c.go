package main

import "fmt"

/*
*
 */
func minTime(skill []int, mana []int) int64 {
	n := len(skill)
	m := len(mana)

	t1 := make([]int, n+1)

	for j := 1; j <= n; j++ {
		t1[j] = t1[j-1] + skill[j-1]*mana[0]
	}

	for i := 1; i < m; i++ {
		t2 := make([]int, n+1)
		t2[1] = t1[1] + skill[0]*mana[i]
		cnt := 0
		for j := 2; j <= n; j++ {
			diff := t2[j-1] + cnt - t1[j]
			if diff < 0 {
				cnt += -diff
			}
			t2[j] = t2[j-1] + skill[j-1]*mana[i]
		}
		for j := 1; j <= n; j++ {
			t2[j] += cnt
		}
		t1 = t2

	}
	return int64(t1[n])

}
func main() {
	fmt.Println(minTime([]int{1, 5, 2, 4}, []int{5, 1, 4, 2}))
}
