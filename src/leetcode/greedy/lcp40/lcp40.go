package main

import (
	"math"
	"sort"
)

/**
力扣挑战赛」心算项目的挑战比赛中，要求选手从 N 张卡牌中选出 cnt 张卡牌，
若这 cnt 张卡牌数字总和为偶数，则选手成绩「有效」且得分为 cnt 张卡牌数字总和。
给定数组 cards 和 cnt，其中 cards[i] 表示第 i 张卡牌上的数字。 请帮参赛选手计算最大的有效得分。若不存在获取有效得分的卡牌方案，则返回 0。

示例 1：

输入：cards = [1,2,8,9], cnt = 3

输出：18

解释：选择数字为 1、8、9 的这三张卡牌，此时可获得最大的有效得分 1+8+9=18。
*/

func maxmiumScore(cards []int, cnt int) int {

	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})

	if cnt == 1 {

		for _, v := range cards {

			if v%2 == 0 {
				return v
			}
		}
	}

	s := 0
	//前cnt个数中最后一个奇数
	odd := 0
	//前cnt个数中最后一个偶数
	lastEven := 0
	//先选前cnt个数
	for _, v := range cards[:cnt] {
		s += v
		if v&1 != 0 {
			odd = v
		} else {
			lastEven = v
		}
	}
	//前n个数和是偶数
	if s%2 == 0 {
		return s
	}

	if len(cards) == cnt {
		return 0
	}

	maxOdd := math.MinInt
	for _, v := range cards[cnt:] {
		if v&1 == 1 {
			maxOdd = v
			break
		}
	}
	maxEven := 0
	for _, v := range cards[cnt:] {
		if v&1 == 0 {
			maxEven = v
			break
		}
	}

	//全为奇数直接返回0
	if lastEven == 0 && maxEven == 0 {
		return 0
	}

	//前n个数中没有偶数
	if maxEven == 0 {
		return s - lastEven + maxOdd
	}
	//后面的数没有奇数
	if maxOdd == math.MinInt {
		return s - odd + lastEven
	}
	//有偶数有奇数
	if maxOdd+odd > lastEven+maxEven && lastEven > 0 {
		return s - lastEven + maxOdd
	}

	return s - odd + maxEven
}

func main() {
	//fmt.Println(maxmiumScore([]int{1, 2, 8, 9, 3, 4, 5, 6, 294, 495, 4, 4, 5}, 13))
	//fmt.Println(maxmiumScore([]int{1, 2, 8, 9, 3, 4, 5, 6, 294, 495, 4, 4, 5, 13}, 14))
	//fmt.Println(maxmiumScore([]int{3, 3, 1}, 1))
	//fmt.Println(maxmiumScore([]int{7, 4, 1}, 1))
	//fmt.Println(maxmiumScore([]int{4, 2, 7, 2, 9, 7, 5, 6}, 3))
}
