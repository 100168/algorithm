package main

import "fmt"

func minArraySum(nums []int, k int, op1 int, op2 int) int {
	n := len(nums)
	f := make([][][]int, n)
	for i := range f {
		f[i] = make([][]int, op1+1)
		for j := range f[i] {
			f[i][j] = make([]int, op2+1)
			for c := range f[i][j] {
				f[i][j][c] = -1
			}
		}
	}
	var dfs func(int, int, int) int

	dfs = func(i, o1, o2 int) int {

		if i < 0 {
			return 0
		}
		if f[i][o1][o2] != -1 {
			return f[i][o1][o2]
		}

		cur := dfs(i-1, o1, o2) + nums[i]

		if o1 > 0 {
			cur = min(cur, dfs(i-1, o1-1, o2)+(nums[i]-1)/2+1)
		}
		if o2 > 0 && nums[i] >= k {
			cur = min(cur, dfs(i-1, o1, o2-1)+nums[i]-k)
		}
		if o1 > 0 && o2 > 0 && nums[i] >= k {
			cur = min(cur, dfs(i-1, o1-1, o2-1)+(nums[i]-k-1)/2+1)

			if (nums[i]-1)/2-k >= 0 {
				cur = min(cur, dfs(i-1, o1-1, o2-1)+((nums[i]-1)/2+1)-k)
			}

		}

		f[i][o1][o2] = cur

		return cur

	}

	return dfs(n-1, op1, op2)

}

func main() {

	fmt.Println(minArraySum([]int{2, 8, 3, 19, 3}, 3, 1, 1))
}
