package main

func smallestSubarrays(nums []int) []int {

	type pair struct{ index, val int }
	n := len(nums)
	suffix := make([]pair, 0)
	ans := make([]int, n)

	suffix = append(suffix, pair{n - 1, nums[n-1]})
	for i := n - 1; i >= 0; i-- {
		temp := suffix
		suffix = nil
		suffix = append(suffix, pair{i, nums[i]})
		maxVal := nums[i] | temp[len(temp)-1].val
		for _, v := range temp {
			if suffix[len(suffix)-1].val == maxVal {
				break
			}
			if v.val|nums[i] == suffix[len(suffix)-1].val {
				continue
			}
			suffix = append(suffix, pair{v.index, v.val | nums[i]})
		}
		cur := suffix[len(suffix)-1]
		if cur.val == maxVal {
			ans[i] = cur.index - i + 1
		}
	}
	return ans
}

func main() {
	println(smallestSubarrays([]int{1, 2}))
}
