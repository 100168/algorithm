package main

import "fmt"

/**
给你一个整数 n。

我们将 1 到 n 的整数按从左到右的顺序排成一个序列。然后，交替地执行以下两种操作，直到只剩下一个整数为止，从操作 1 开始：

操作 1：从左侧开始，隔一个数删除一个数。
操作 2：从右侧开始，隔一个数删除一个数。
返回最后剩下的那个整数。



示例 1：

输入： n = 8

输出： 3

解释：

写下序列 [1, 2, 3, 4, 5, 6, 7, 8]。
从左侧开始，我们删除每隔一个数字：[1, 2, 3, 4, 5, 6, 7, 8]。剩下的整数是 [1, 3, 5, 7]。
从右侧开始，我们删除每隔一个数字：[1, 3, 5, 7]。剩下的整数是 [3, 7]。
从左侧开始，我们删除每隔一个数字：[3, 7]。剩下的整数是 [3]。
示例 2：

输入： n = 5

输出： 1

解释：

写下序列 [1, 2, 3, 4, 5]。
从左侧开始，我们删除每隔一个数字：[1, 2, 3, 4, 5]。剩下的整数是 [1, 3, 5]。
从右侧开始，我们删除每隔一个数字：[1, 3, 5]。剩下的整数是 [1, 5]。
从左侧开始，我们删除每隔一个数字：[1, 5]。剩下的整数是 [1]。
示例 3：

输入： n = 1

输出： 1

解释：

写下序列 [1]。
最后剩下的整数是 1。

提示：

1 <= n <= 1015
*/

func lastInteger(n int64) int64 {
	start, d := int64(1), int64(1) // 等差数列的首项和公差
	for ; n > 1; n = (n + 1) / 2 {
		start += (n - 2 + n%2) * d
		d *= -2
	}
	return start
}

func main() {
	// 测试用例1: n = 8, 期望输出 3
	result1 := lastInteger(8)
	fmt.Printf("n = 8, 结果: %d, 期望: 3", result1)
	if result1 == 3 {
		fmt.Println(" ✓")
	} else {
		fmt.Println(" ✗")
	}

	// 测试用例2: n = 5, 期望输出 1
	result2 := lastInteger(5)
	fmt.Printf("n = 5, 结果: %d, 期望: 1", result2)
	if result2 == 1 {
		fmt.Println(" ✓")
	} else {
		fmt.Println(" ✗")
	}

	// 测试用例3: n = 1, 期望输出 1
	result3 := lastInteger(1)
	fmt.Printf("n = 1, 结果: %d, 期望: 1", result3)
	if result3 == 1 {
		fmt.Println(" ✓")
	} else {
		fmt.Println(" ✗")
	}

	// 测试用例4: n = 10^15 (大数值测试，验证不会内存溢出)
	largeN := int64(1000000000000000) // 10^15
	result4 := lastInteger(int64(int(largeN)))
	fmt.Printf("n = %d, 结果: %d (大数值测试通过，无内存溢出)", largeN, result4)
	fmt.Println(" ✓")
}
