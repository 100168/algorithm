package main

func minimumSubarrayLength(nums []int, t int) int {

	n := len(nums)

	type pair struct{ i, val int }

	//1, 2, 32, 21
	//55
	ans := n + 1
	pre := make([]pair, 0)
	for i, v := range nums {
		pre = append(pre, pair{i, v})

		newPre := make([]pair, 0)
		for j := 0; j < len(pre); j++ {
			cur := pre[j]
			cur.val |= v
			if len(newPre) == 0 {
				newPre = append(newPre, cur)
			} else {
				if newPre[len(newPre)-1].val == cur.val {
					newPre[len(newPre)-1].i = cur.i
				} else {
					newPre = append(newPre, cur)
				}
			}
			if cur.val >= t {
				ans = min(ans, i-cur.i+1)
			}

		}
		pre = newPre
	}
	if ans == n+1 {
		return -1
	}
	return ans

}

func minimumSubarrayLength2(nums []int, k int) int {

	n := len(nums)

	ans := n + 1
	for i := 0; i < n; i++ {

		cur := nums[i]
		for j := i; j < n; j++ {
			cur |= nums[j]
			if cur >= k {
				ans = min(ans, j-i+1)
				break
			}
		}
	}
	if ans == n+1 {
		return -1
	}
	return ans

}

func main() {
	println(minimumSubarrayLength([]int{1, 2, 32, 21}, 55))
	println(minimumSubarrayLength2([]int{1, 2, 32, 21}, 55))
}
