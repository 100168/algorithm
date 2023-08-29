package main

import "fmt"

type BitTree struct {
	sum    []int
	length int
}

func NewBitTree(n int) *BitTree {
	return &BitTree{
		sum:    make([]int, n+1),
		length: n + 1,
	}
}

func (b *BitTree) lowBit(index int) int {
	return index & -index
}

func (b *BitTree) query(index int) int {

	ans := 0
	for ; index < b.length && index > 0; index -= b.lowBit(index) {
		ans += b.sum[index]
	}
	return ans
}

func (b *BitTree) update(index int, value int) {
	for ; index < b.length; index += b.lowBit(index) {
		b.sum[index] += value
	}
}

func main() {
	s := "purchase_apply_order_status_log"
	fmt.Println(len(s))
}
