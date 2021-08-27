package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有next 指针都被设置为 NULL。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func connect(root *Node) *Node {

	if root == nil {
		return nil
	}
	resNode := root
	curren := root

	for curren != nil {
		nextNode := curren
		for nextNode != nil {
			if nextNode.Left != nil {
				nextNode.Left.Next = nextNode.Right
				if nextNode.Next != nil {
					nextNode.Right.Next = nextNode.Next.Left
				}
			}
			nextNode = nextNode.Next
		}
		curren = curren.Left
	}

	return resNode
}
func main() {

}
