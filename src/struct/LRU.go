package main

type LRUCache struct {
	nodeMap map[int]*CacheNode
	head    *CacheNode
	tail    *CacheNode
	size    int
}

type CacheNode struct {
	Pre  *CacheNode
	Next *CacheNode
	Val  int
	Key  int
}

func Constructor(capacity int) LRUCache {
	cache := new(LRUCache)
	cache.size = capacity
	cache.head = new(CacheNode)
	cache.tail = new(CacheNode)
	cache.head.Next = cache.tail
	cache.tail.Pre = cache.head
	cache.nodeMap = make(map[int]*CacheNode)
	return *cache
}

func (cache *LRUCache) Get(key int) int {
	nodeMap := cache.nodeMap
	if nodeMap[key] == nil {
		return -1
	}
	cache.moveToHead(nodeMap[key])
	return nodeMap[key].Val
}
func (cache *LRUCache) moveToHead(node *CacheNode) {
	//将当前节点断开
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	//移动到头节点
	headNext := cache.head.Next
	cache.head.Next = node
	node.Pre = cache.head
	node.Next = headNext
	headNext.Pre = node
}

func (cache *LRUCache) removeTail() {
	delete(cache.nodeMap, cache.tail.Pre.Key)
	cache.tail.Pre = cache.tail.Pre.Pre
	cache.tail.Pre.Next = cache.tail

}
func (cache *LRUCache) addTail(node *CacheNode) {
	tailPre := cache.tail.Pre
	cache.tail.Pre = node
	node.Next = cache.tail
	node.Pre = tailPre
	tailPre.Next = node
}

func (cache *LRUCache) Put(key int, value int) {
	nodeMap := cache.nodeMap
	curNode := nodeMap[key]
	if curNode == nil {
		curNode = new(CacheNode)
		cache.addTail(curNode)
	}
	curNode.Val = value
	curNode.Key = key
	cache.moveToHead(curNode)
	nodeMap[key] = curNode
	if len(nodeMap) > cache.size {
		cache.removeTail()
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 0)
	cache.Put(2, 2)
	println(cache.Get(1))
	cache.Put(3, 3)
	println(cache.Get(2))
	cache.Put(4, 4)
	println(cache.Get(1))
	println(cache.Get(3))
	println(cache.Get(4))

}
