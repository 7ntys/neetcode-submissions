type LRUCache struct {
    cap int
	cache map[int]*Node
	left *Node
	right *Node
}

type Node struct {
	key int
	val int
	prev *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
    left, right := &Node{}, &Node{}
	left.next, right.prev = right, left
	cache := make(map[int]*Node)
	return LRUCache{cap: capacity, cache: cache, left: left, right: right}
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

func (this *LRUCache) add(node *Node) {
	prev := this.right.prev
	prev.next, this.right.prev = node, node
	node.prev = prev
	node.next = this.right
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.cache[key]; ok {
		this.remove(n)
		this.add(n)
		return n.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		n := this.cache[key]
		this.remove(n)
		n.val = value
		this.add(n)
	} else {
		n := &Node{val: value, key: key}
		this.add(n)
		this.cache[key] = n
	}

	if len(this.cache) > this.cap {
		lru := this.left.next
		delete(this.cache, lru.key)
		this.remove(lru)
	}
}
