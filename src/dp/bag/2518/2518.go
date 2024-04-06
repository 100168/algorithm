package main

import "strconv"

func countPartitions(nums []int, k int) int {

	n := len(nums)
	mod := 1_000_000_007

	sum := 0

	for i := range nums {
		sum += nums[i]
	}
	cache := make(map[string]int)

	var dfs func(int, int) int
	dfs = func(i int, rest int) int {

		if i >= n {
			if rest >= k && sum-rest >= k {
				return 1
			}
			return 0
		}
		key := strconv.Itoa(i) + "-" + strconv.Itoa(rest)
		if _, ok := cache[key]; ok {
			return cache[key]
		}
		cur := (dfs(i+1, rest)) + (dfs(i+1, rest+nums[i]))
		cur %= mod
		cache[key] = cur
		return cur

	}
	return dfs(0, 0)
}

func main() {
	println(countPartitions([]int{977208288, 291246471, 396289084, 732660386, 353072667, 34663752, 815193508, 717830630, 566248717, 260280127, 824313248, 701810861, 923747990, 478854232, 781012117, 525524820, 816579805, 861362222, 854099903, 300587204, 746393859, 34127045, 823962434, 587009583, 562784266, 115917238, 763768139, 393348369, 3433689, 586722616, 736284943, 596503829, 205828197, 500187252, 86545000, 490597209, 497434538, 398468724, 267376069, 514045919, 172592777, 469713137, 294042883, 985724156, 388968179, 819754989, 271627185, 378316864, 820060916, 436058499, 385836880, 818060440, 727928431, 737435034, 888699172, 961120185, 907997012, 619204728, 804452206, 108201344, 986517084, 650443054}, 95))
}
