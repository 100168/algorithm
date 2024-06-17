package main

import "fmt"

func countOfPeaks(nums []int, queries [][]int) []int {

	bt := new(bitTree)

	bt.sum = make([]int, len(nums)+1)
	bt.n = len(nums) + 1

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			bt.update(i, 1)
		}
	}
	ans := make([]int, 0)
	for i := range queries {
		op, l, r := queries[i][0], queries[i][1], queries[i][2]
		if op == 1 {
			cur := bt.query(r-1) - bt.query(l)
			ans = append(ans, max(cur, 0))
		} else {
			if l > 0 && l+1 < len(nums) {
				flag := false
				if r > nums[l-1] && r > nums[l+1] {
					flag = true
				}
				if nums[l] > nums[l-1] && nums[l] > nums[l+1] {
					if flag == false {
						bt.update(l, -1)
					}
				} else {
					if flag == true {
						bt.update(l, +1)
					}
				}
			}

			if l+2 < len(nums) {
				flag := false
				if nums[l+1] > r && nums[l+1] > nums[l+2] {
					flag = true
				}
				if nums[l+1] > nums[l] && nums[l+1] > nums[l+2] {
					if flag == false {
						bt.update(l+1, -1)
					}
				} else {
					if flag == true {
						bt.update(l+1, +1)
					}
				}
			}

			if l > 1 {
				flag := false
				if nums[l-1] > r && nums[l-1] > nums[l-2] {
					flag = true
				}
				if nums[l-1] > nums[l] && nums[l-1] > nums[l-2] {
					if flag == false {
						bt.update(l-1, -1)
					}
				} else {
					if flag == true {
						bt.update(l-1, +1)
					}
				}
			}
			nums[l] = r

		}
	}
	return ans
}

type bitTree struct {
	sum []int
	n   int
}

func lowBit(index int) int {
	return index & -index
}

func (bt bitTree) query(index int) int {
	ans := 0
	for index > 0 {
		ans += bt.sum[index]
		index -= lowBit(index)

	}
	return ans
}

func (bt bitTree) update(index int, value int) {
	for index < bt.n {
		bt.sum[index] += value
		index += lowBit(index)
	}
}

func main() {
	fmt.Println(countOfPeaks([]int{8, 7, 10}, [][]int{{2, 2, 4}, {2, 1, 9}, {1, 0, 2}}))
}
