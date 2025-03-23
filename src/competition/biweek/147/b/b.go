package main

import "github.com/emirpasic/gods/trees/redblacktree"

type TaskManager struct {
	rbt      *redblacktree.Tree
	indexMap map[int]int

	priorityMap map[int]int
}

type pair struct {
	priority int
	taskId   int
}

func Constructor(tasks [][]int) TaskManager {

	ts := new(TaskManager)
	ts.indexMap = make(map[int]int)
	ts.priorityMap = make(map[int]int)

	cmp := func(a, b any) int {

		x, y := a.(pair), b.(pair)
		if x.priority == y.priority {
			return y.taskId - x.taskId
		}
		return y.priority - x.priority
	}

	ts.rbt = redblacktree.NewWith(cmp)

	for _, v := range tasks {
		u, t, p := v[0], v[1], v[2]

		ts.indexMap[t] = u
		ts.priorityMap[t] = p
		ts.rbt.Put(pair{p, t}, u)
	}

	return *ts
}

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
	tm.rbt.Put(pair{priority, taskId}, userId)
	tm.priorityMap[taskId] = priority
	tm.indexMap[taskId] = userId
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {

	userId := tm.indexMap[taskId]

	prP := tm.priorityMap[taskId]
	tm.rbt.Remove(pair{prP, taskId})

	tm.priorityMap[taskId] = newPriority
	tm.rbt.Put(pair{newPriority, taskId}, userId)
}

func (tm *TaskManager) Rmv(taskId int) {

	prP := tm.priorityMap[taskId]
	tm.rbt.Remove(pair{prP, taskId})
	delete(tm.indexMap, taskId)
}

func (tm *TaskManager) ExecTop() int {

	left := tm.rbt.Left()

	if left == nil {
		return -1
	}

	tm.rbt.Remove(left.Key)

	c := left.Value.(int)
	return c

}
