package main

/*
任务调度优化是计算机性能优化的关键任务之一。在任务众多时，不同的调度策略可能会得到不同的总体执行时间，因此寻求一个最优的调度方案是非常有必要的。

通常任务之间是存在依赖关系的，即对于某个任务，你需要先完成他的前导任务（如果非空），才能开始执行该任务。
我们保证任务的依赖关系是一棵二叉树，其中 root 为根任务，root.left 和 root.right 为他的两个前导任务（可能为空），root.val 为其自身的执行时间。

在一个 CPU 核执行某个任务时，我们可以在任何时刻暂停当前任务的执行，并保留当前执行进度。
在下次继续执行该任务时，会从之前停留的进度开始继续执行。暂停的时间可以不是整数。

现在，系统有两个 CPU 核，即我们可以同时执行两个任务，但是同一个任务不能同时在两个核上执行。给定这颗任务树，请求出所有任务执行完毕的最小时间。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
设一颗任务树的左子树所有任务时间和为 a ，最大并行时间为 b ，右子树分别为 c,d 。那么这颗任务树最大并行时间为a+c，但不一定能取到。
不失一般性，假设 a≥c 。如果 a−2b≤c ，那么最优策略是左树先并行 a−c 的任务，剩余 c 的任务量，然后再两树并行，就达成了一直并行。并行时间取到了
a+c

	。

如果 a−2b>c ，那么剩余 a−2b−c 的任务无法被并行。并行时间最大为 b+c 。
*/
func minimalExecTime(root *TreeNode) float64 {

	var dfs func(*TreeNode) (float64, float64)

	dfs = func(node *TreeNode) (float64, float64) {
		if node == nil {
			return 0.0, 0.0
		}
		tc := float64(node.Val)
		a, b := dfs(node.Left)
		c, d := dfs(node.Right)
		tc = tc + a + c
		pc := float64(node.Val) + max((a+c)/2, b, d)
		return tc, pc

	}
	_, pc := dfs(root)
	return pc
}

/*
*
方法二有点SB
*/
func minimalExecTime2(root *TreeNode) float64 {

	var dfs func(*TreeNode) (float64, float64)

	dfs = func(node *TreeNode) (float64, float64) {
		if node == nil {
			return 0.0, 0.0
		}
		tc := float64(node.Val)
		a, b := dfs(node.Left)
		c, d := dfs(node.Right)
		tc = tc + a + c
		if a > c {
			a, c = c, a
			b, d = d, b
		}
		pc := float64(0)
		if a-2*b < c {
			pc = (a + c) / 2
		} else {
			pc = c + b
		}

		return tc, pc

	}
	tc, pc := dfs(root)
	return tc - pc
}
