package main

type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*LinkNode
	head     *LinkNode
	tail     *LinkNode
}

type LinkNode struct {
	k    int
	v    int
	next *LinkNode
	pre  *LinkNode
}

func initNode(k, v int) *LinkNode {
	return &LinkNode{
		k:    k,
		v:    v,
		next: nil,
		pre:  nil,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*LinkNode{},
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; ok {
		node := this.cache[key]
		this.moveToHead(node)
		return node.v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		node := this.cache[key]
		node.v = value
		this.moveToHead(node)
	} else {
		node := initNode(key, value)
		this.cache[key] = node
		this.toHead(node)
		this.size++
		if this.size > this.capacity {
			n := this.removeTail()
			delete(this.cache, n.k)
			this.size--
		}
	}
}
func (l *LRUCache) removeNode(node *LinkNode) {
	node.next.pre = node.pre
	node.pre.next = node.next
}
func (l *LRUCache) removeTail() *LinkNode {
	node := l.tail.pre
	l.removeNode(node)
	return node
}
func (l *LRUCache) toHead(node *LinkNode) {
	node.next = l.head.next
	node.pre = l.head
	l.head.next.pre = node
	l.head.next = node
}

func (l *LRUCache) moveToHead(node *LinkNode) {
	l.removeNode(node)
	l.toHead(node)
}
