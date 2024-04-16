package main

func beautifulSubarrays(nums []int) (ans int64) {
	s := make([]int, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] ^ x
	}
	cnt := map[int]int{}
	for _, x := range s {
		// 先计入答案再统计个数，如果反过来的话，就相当于把空子数组也计入答案了
		ans += int64(cnt[x])
		cnt[x]++
	}
	return
}

func main() {
	println(beautifulSubarrays([]int{0, 1, 2, 3}))
}
