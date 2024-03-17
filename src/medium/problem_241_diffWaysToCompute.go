package main

/*输入: "2*3-4*5"
输出: [-34, -14, -10, -10, 10]
解释:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/different-ways-to-add-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
import (
	"strconv"
)

func diffWaysToCompute(input string) []int {
	// 如果是数字，直接返回
	if isDigit(input) {
		tmp, _ := strconv.Atoi(input)
		return []int{tmp}
	}

	// 空切片
	var res []int
	// 遍历字符串
	for index, c := range input {
		tmpC := string(c)
		if tmpC == "+" || tmpC == "-" || tmpC == "*" {
			// 如果是运算符，则计算左右两边的算式
			left := diffWaysToCompute(input[:index])
			right := diffWaysToCompute(input[index+1:])

			for _, leftNum := range left {
				for _, rightNum := range right {
					var addNum int
					if tmpC == "+" {
						addNum = leftNum + rightNum
					} else if tmpC == "-" {
						addNum = leftNum - rightNum
					} else {
						addNum = leftNum * rightNum
					}
					res = append(res, addNum)
				}
			}
		}
	}

	return res
}

// 判断是否为全数字
func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return true
}

// dfs
func diffWaysToCompute2(expression string) []int {
	// 记忆化
	subExpWays := make(map[string][]int)

	// 利用函数闭包, 避免一些不必要的传值
	var helper func(expression string) []int
	helper = func(expression string) []int {
		ways := make([]int, 0)

		for i := range expression {
			ch := expression[i]
			if ch == '-' || ch == '+' || ch == '*' {
				var (
					leftExp   = expression[:i]
					rightExp  = expression[i+1:]
					leftWays  []int
					rightWays []int
					ok        bool
				)
				// 先检查 map
				if leftWays, ok = subExpWays[leftExp]; !ok {
					leftWays = helper(leftExp)
				}
				if rightWays, ok = subExpWays[rightExp]; !ok {
					rightWays = helper(rightExp)
				}

				for _, l := range leftWays {
					for _, r := range rightWays {
						switch ch {
						case '-':
							ways = append(ways, l-r)
						case '+':
							ways = append(ways, l+r)
						case '*':
							ways = append(ways, l*r)
						}
					}
				}
			}
		}
		if len(ways) == 0 {
			num, _ := strconv.Atoi(expression)
			ways = append(ways, num)
		}

		subExpWays[expression] = ways
		return ways
	}

	ways := helper(expression)

	return ways
}

func diffWaysToCompute3(expression string) []int {
	m := len(expression)
	nums := make([]int, 0, m/4)
	ops := make([]byte, 0, m/2)

	// 根据操作符划分数字，统计数字和操作符（数字个数 = 操作符个数 + 1），
	for i := 0; i < m; i++ {
		if expression[i] < '0' || expression[i] > '9' {
			ops = append(ops, expression[i])
		} else {
			num := 0
			for i < m && expression[i] >= '0' && expression[i] <= '9' {
				num = num*10 + int(expression[i]-'0')
				i++
			}
			nums = append(nums, num)
			i--
		}
	}

	n := len(nums)
	dp := make([][][]int, n)

	// dp[j][i][k] 表示对于数字 j 到数字 i 部分的子表达式可以有 dp[i][j][k] 种结果（结果可重复）
	for i := range dp {
		dp[i] = make([][]int, n)

		for j := i; j >= 0; j-- {
			// i == j，数字 j 到 i 就是一个数字
			if i == j {
				dp[j][i] = append(dp[j][i], nums[i])
				continue
			}

			// 遍历从 j 到 i 的每种可能结果
			for k := j; k < i; k++ {
				// 从 j 到 k (j <= k < i) 的子表达式可能的结果，
				// 即当前操作符 ops[k] 左边的子表达式
				for _, left := range dp[j][k] {
					// 从 k 到 i 的子表达式可能的结果，
					// 即当前操作符 ops[k] 右边的子表达式
					for _, right := range dp[k+1][i] {
						var val int
						switch ops[k] {
						case '-':
							val = left - right
						case '+':
							val = left + right
						case '*':
							val = left * right
						}
						dp[j][i] = append(dp[j][i], val)
					}
				}
			}
		}
	}

	return dp[0][n-1]
}
