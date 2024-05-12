package main

/**
给你一个下标从 0 开始的整数数组 nums 。nums 的一个子数组如果满足以下条件，那么它是 不间断 的：

i，i + 1 ，...，j  表示子数组中的下标。对于所有满足 i <= i1, i2 <= j 的下标对，都有 0 <= |nums[i1] - nums[i2]| <= 2 。
请你返回 不间断 子数组的总数目。

子数组是一个数组中一段连续 非空 的元素序列。*/

func continuousSubarrays(nums []int) int64 {

	ans := 0
	l := 0

	decreaseSt := make([]int, 0)
	increaseSt := make([]int, 0)
	for r, v := range nums {
		for len(increaseSt) > 0 && nums[increaseSt[len(increaseSt)-1]] <= v {
			increaseSt = increaseSt[:len(increaseSt)-1]
		}
		for len(decreaseSt) > 0 && nums[decreaseSt[len(decreaseSt)-1]] >= v {
			decreaseSt = decreaseSt[:len(decreaseSt)-1]
		}
		increaseSt = append(increaseSt, r)
		decreaseSt = append(decreaseSt, r)
		for len(decreaseSt) > 0 && abs(nums[decreaseSt[0]]-nums[increaseSt[0]]) > 2 {
			if decreaseSt[0] == l {
				decreaseSt = decreaseSt[1:]
			}
			if increaseSt[0] == l {
				increaseSt = increaseSt[1:]
			}
			l++
		}

		ans += r - l + 1
	}
	return int64(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func continuousSubarrays2(a []int) (ans int64) {
	cnt := map[int]int{}
	left := 0
	for right, x := range a {
		cnt[x]++
		for {
			mx, mn := x, x
			for k := range cnt {
				mx = max(mx, k)
				mn = min(mn, k)
			}
			if mx-mn <= 2 {
				break
			}
			y := a[left]
			if cnt[y]--; cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans += int64(right - left + 1)
	}
	return
}

func main() {
	println(continuousSubarrays([]int{5, 6, 7, 5, 5, 4, 4}))
}
