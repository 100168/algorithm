package main

func minOperations2(nums []int) int {

	cnt := make(map[int]int)

	for i := range nums {
		cnt[nums[i]]++
	}
	ans := 0
	for i := range cnt {
		if cnt[i] == 1 {
			return -1
		}
		rest := cnt[i] % 3
		mul := cnt[i] / 3
		ans += mul
		if rest == 2 || rest == 1 {
			ans++
		}

	}
	return ans
}

func main() {
	minOperations2([]int{2, 3, 3, 2, 2, 4, 2, 3, 4})
}
