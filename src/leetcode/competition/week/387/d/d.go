package main

import "slices"

func resultArray(nums []int) []int {

	n := len(nums)
	sa := make([]int, n)
	for i := range sa {
		sa[i] = i
	}
	rk := make([]int, n)

	slices.SortStableFunc(sa, func(a, b int) int {
		return nums[a] - nums[b]
	})
	for i := range rk {
		rk[sa[i]] = i
	}
	a := make([]int, 0)
	b := make([]int, 0)

	a = append(a, nums[0])
	b = append(b, nums[1])
	b1 := new(BitSet)
	b1.sum = make([]int, n+1)
	b1.len = n + 1
	b2 := new(BitSet)
	b2.sum = make([]int, n+1)
	b2.len = n + 1
	b1.update(rk[0] + 1)
	b2.update(rk[1] + 1)
	//5,14,3,1,2
	for i := 2; i < n; i++ {

		cntA := b1.query(n) - b1.query(rk[i]+1)
		cntB := b2.query(n) - b2.query(rk[i]+1)
		if cntA > cntB {
			a = append(a, nums[i])
			b1.update(rk[i] + 1)
		} else if cntA < cntB {
			b = append(b, nums[i])
			b2.update(rk[i] + 1)
		} else {
			if len(a) <= len(b) {
				a = append(a, nums[i])
				b1.update(rk[i] + 1)
			} else {
				b = append(b, nums[i])
				b2.update(rk[i] + 1)
			}
		}
	}

	ans := make([]int, n)
	index := 0
	for i := range a {
		ans[index] = a[i]
		index++
	}
	for i := range b {
		ans[index] = b[i]
		index++
	}

	return ans
}

type BitSet struct {
	sum []int
	len int
}

func (b *BitSet) lowBit(index int) int {

	return index & -index
}

func (b *BitSet) query(index int) int {

	ans := 0
	for index > 0 {
		ans += b.sum[index]
		index -= b.lowBit(index)
	}
	return ans
}
func (b *BitSet) update(index int) {

	for index < b.len {
		b.sum[index] += 1
		index += b.lowBit(index)
	}
}
func main() {

	print(resultArray([]int{5, 14, 3, 1, 2}))
}
