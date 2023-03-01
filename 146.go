package main

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

type DLinkNode struct {
	k, v       int
	prev, next *DLinkNode
}

func initDLinkedNode(key, value int) *DLinkNode {
	return &DLinkNode{
		k: key,
		v: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DLinkNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	res := this.cache[key]
	this.moveToHead(res)
	return res.v
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			n := this.removeTail()
			delete(this.cache, n.k)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.v = value
		this.moveToHead(node)
	}
}

func (l *LRUCache) addToHead(node *DLinkNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeNode(node *DLinkNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (l *LRUCache) removeTail() *DLinkNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}

func (l *LRUCache) moveToHead(node *DLinkNode) {
	l.removeNode(node)
	l.addToHead(node)
}
