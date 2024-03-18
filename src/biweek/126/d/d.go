package main

type bitSet struct {
	bit []int
	n   int
}

func lowbit(index int) int {
	return index & -index
}
func (b *bitSet) query(index int) int {

	ans := 0
	for ; index > 0; index -= lowbit(index) {
		ans += b.bit[index]
	}
	return ans
}

func (b *bitSet) update(index int, val int) {

	for index < b.n {
		b.bit[index] += val
		index += lowbit(index)
	}
}

type NumArray struct {
	b *bitSet
}

func Constructor(nums []int) NumArray {

	na := new(NumArray)
	bs := new(bitSet)
	n := len(nums)
	bs.n = n + 1
	bs.bit = make([]int, n+1)
	for i := 1; i <= n; i++ {
		bs.update(i, nums[i-1])
	}
	na.b = bs
	return *na
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.b.query(right+1) - this.b.query(left)
}
