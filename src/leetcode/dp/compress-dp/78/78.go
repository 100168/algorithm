package main

func subsets(nums []int) [][]int {
	n := len(nums)

	ans := make([][]int, 0)
	for i := 0; i < 1<<n; i++ {
		cur := make([]int, 0)
		for j := 0; j < n; j++ {
			if 1<<j&i != 0 {
				cur = append(cur, nums[j])
			}
		}
		ans = append(ans, cur)
	}
	return ans
}

func subsets2(nums []int) [][]int {

	n := len(nums)

	ans := make([][]int, 0)

	mask := 1<<n - 1
	next := mask
	for next > 0 {
		cur := make([]int, 0)
		for i := 0; i < n; i++ {
			if 1<<i&next != 0 {
				cur = append(cur, nums[i])
			}

		}
		next = (next - 1) & mask
		ans = append(ans, cur)
	}
	ans = append(ans, []int{})
	return ans
}
