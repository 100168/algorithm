package main

import (
	"sort"
)

func numberOfSubarrays(nums []int) int64 {
	n := len(nums)
	ans := int64(0)
	sum := make([]int, n)
	b := new(bitSet)
	indexMap := make(map[int][]int)
	for i, v := range nums {
		indexMap[v] = append(indexMap[v], i)
	}
	newNums := make([]int, 0)
	for k := range indexMap {
		newNums = append(newNums, k)
	}
	sort.Ints(newNums)

	rkMap := make(map[int]int)
	for i, v := range newNums {
		rkMap[v] = i + 1
	}
	b.sum = make([]int, len(rkMap)+1)
	b.len = len(rkMap) + 1

	for i, v := range nums {
		rk := rkMap[v]
		q := b.query(rk)
		sum[i] = q
		b.update(rk)
	}
	for _, v := range indexMap {
		left := 0
		for i := 0; i < len(v); i++ {
			right := v[i]
			leftV := sum[v[left]]
			rightV := sum[right]
			diff := rightV - leftV
			if diff == right-v[left] {
				ans += int64(i - left + 1)
			} else {
				left = i
				i--
			}
		}
	}
	return ans

}

type bitSet struct {
	sum []int
	len int
}

func lowBit(index int) int {

	return index & -index
}

func (b bitSet) query(index int) int {
	ans := 0
	for index > 0 {
		ans += b.sum[index]
		index -= lowBit(index)
	}
	return ans
}

func (b bitSet) update(index int) {
	for index < b.len {
		b.sum[index]++
		index += lowBit(index)
	}
}

func main() {

	println(numberOfSubarrays([]int{1, 4, 3, 3, 2, 4, 4}))
}
