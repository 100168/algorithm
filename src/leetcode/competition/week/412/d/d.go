package main

import (
	"fmt"
	"sort"
	"strconv"
)

func countPairs(nums []int) int {

	sort.Ints(nums)
	ans := 0

	for i, v1 := range nums {

	next:
		for _, v2 := range nums[i+1:] {

			if v2 == v1 {
				ans++
				continue next
			}
			itoa := strconv.Itoa(v2)

			bytes := []byte(itoa)

			for c := range bytes {
				for cc := c + 1; cc < len(bytes); cc++ {
					bytes[c], bytes[cc] = bytes[cc], bytes[c]
					s := string(bytes)
					atoi, _ := strconv.Atoi(s)

					if atoi == v1 {
						ans++
						continue next
					}
					bytes[c], bytes[cc] = bytes[cc], bytes[c]
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countPairs([]int{5, 12, 8, 5, 5, 1, 20, 3, 10, 10, 5, 5, 5, 5, 1}))
}
