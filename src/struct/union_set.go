package main

type UnionSet struct {
	parent map[int]int
	size   map[int]int
}

func (unionSet *UnionSet) findFather(cur int) int {
	var stack []int
	for cur != unionSet.parent[cur] {
		cur = unionSet.parent[cur]
		stack = append(stack, cur)
	}

	for len(stack) > 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		unionSet.parent[pop] = cur
	}
	return cur
}

func (unionSet *UnionSet) isSameSet(v1 int, v2 int) bool {
	return unionSet.findFather(v1) == unionSet.findFather(v2)
}

func (unionSet *UnionSet) union(v1 int, v2 int) {

	father1 := unionSet.findFather(v1)
	father2 := unionSet.findFather(v2)

	newSize := unionSet.size[father1] + unionSet.size[father2]
	if father1 != father2 {
		if unionSet.size[father1] > unionSet.size[father2] {
			unionSet.parent[father2] = father1
			unionSet.size[father1] = newSize
			delete(unionSet.size, unionSet.size[father2])
		} else {
			unionSet.parent[father1] = father2
			unionSet.size[father2] = newSize
			delete(unionSet.size, unionSet.size[father1])

		}
	}
}
