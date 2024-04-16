package main

import "fmt"

func permuteUnique(nums []int) [][]int {

	var flag uint8
	var ans [][]int
	path := make([]int, 0)
	getAllSort(nums, &ans, flag, path)
	return ans

}

func getAllSort(nums []int, ans *[][]int, flag uint8, path []int) {

	if len(path) == len(nums) {

		//fmt.Println(path)
		*ans = append(*ans, path)
		return
	}
	check := make([]int, 21)

	for i, v := range nums {

		if (flag>>i)&1 == 0 && check[10-v] == 0 {
			flag ^= 1 << i
			check[10-v] = 1
			getAllSort(nums, ans, flag, append(path, v))
			flag ^= 1 << i
		}

	}

}
func main() {

	a := []int{-1, 2, -1, 2, 1, -1, 2, 1}
	unique := permuteUnique(a)

	fmt.Println(unique)
}
