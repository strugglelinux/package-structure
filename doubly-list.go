package structure

import (
	"sync"
	"sync/atomic"
)

//Node 节点
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

//DoublyList 双向列表
type DoublyList struct {
	mu   sync.RWMutex
	size uint64
	head *Node
	tail *Node
}

//NewDoublyList 初始化链表
func NewDoublyList() *DoublyList {
	dl := new(DoublyList)
	dl.size = 0
	dl.head = nil
	dl.tail = nil
	return dl
}

//Append 添加数据
func (dl *DoublyList) Append(data interface{}) bool {
	dl.mu.Lock()
	defer dl.mu.Unlock()

	return dl.insertTail(data)
}

//insertTail 尾部插入
func (dl *DoublyList) insertTail(data interface{}) bool {
	if data == nil {
		return false
	}
	node := new(Node)
	node.data = data
	if dl.size == 0 {
		node.next = nil
		node.prev = nil
		dl.head = node
		dl.tail = node
	} else { //尾部添加节点
		node.prev = dl.tail
		node.next = nil
		//设置尾节点的下一个节点
		dl.tail.next = node
		//设置链表尾节点
		dl.tail = node
	}
	atomic.AddUint64(&dl.size, 1)
	return true
}

//insertBefore 节点之前插入
func (dl *DoublyList) insertBefore(node *Node, data interface{}) bool {
	if node == nil || data == nil {
		return false
	}
	newNode := new(Node)
	newNode.data = data
	newNode.prev = node.prev
	newNode.next = node
	if ok := dl.isHead(node); !ok { //非头节点
		node.prev.next = newNode
	}
	node.prev = newNode
	newNode.next = node
	dl.head = newNode
	atomic.AddUint64(&dl.size, 1)
	return true
}

//insertMiddle 中间插入
func (dl *DoublyList) insertMiddle(node *Node, data interface{}) bool {
	if node == nil || data == nil {
		return false
	}
	newNode := new(Node)
	newNode.data = data
	newNode.prev = node
	newNode.next = node.next
	node.next.prev = newNode
	node.next = newNode
	atomic.AddUint64(&dl.size, 1)
	return true
}

//InsertBefore 插入节点前
func (dl *DoublyList) InsertBefore(node *Node, data interface{}) bool {
	if node == nil || data == nil {
		return false
	}
	dl.mu.Lock()
	defer dl.mu.Unlock()
	return dl.insertBefore(node, data)
}

//InsertNext 插入节点后
func (dl *DoublyList) InsertNext(node *Node, data interface{}) bool {
	if node == nil || data == nil {
		return false
	}
	var result bool
	dl.mu.Lock()
	defer dl.mu.Unlock()
	if dl.isTail(node) { //判断是否为尾节点
		result = dl.insertTail(data)
	} else {
		result = dl.insertMiddle(node, data)
	}
	return result
}

//Remove 移除节点
func (dl *DoublyList) Remove(node *Node) bool {
	if node == nil || dl.size == 0 {
		return false
	}
	dl.mu.Lock()
	defer dl.mu.Unlock()
	//头节点
	if ok := dl.isHead(node); ok {
		dl.head = node.next
		if node.next != nil {
			node.next.prev = nil
			node.next = nil
		}
	} else if ok := dl.isTail(node); ok { //尾节点
		dl.tail = node.prev
		node.prev.next = nil
	} else { //中间节点
		node.prev.next = node.next
		node.next.prev = node.prev
		node.next = nil
		node.prev = nil
	}
	//减1操作
	atomic.AddUint64(&dl.size, ^uint64(0))
	return true
}

//是否为尾节点
func (dl *DoublyList) isTail(node *Node) bool {
	if node == nil {
		return false
	}
	return node.next == nil
}

//是否未头节点
func (dl *DoublyList) isHead(node *Node) bool {
	return node.prev == nil
}

//GetHead 头节点
func (dl *DoublyList) GetHead() *Node {
	dl.mu.RLock()
	defer dl.mu.RUnlock()
	return dl.head
}

//GetTail 获取尾节点
func (dl *DoublyList) GetTail() *Node {
	dl.mu.RLock()
	defer dl.mu.RUnlock()
	return dl.tail
}

//Clear 清空链表
func (dl *DoublyList) Clear() {
	dl.size = 0
	dl.head = nil
	dl.tail = nil
}

//Size 节点数量
func (dl *DoublyList) Size() uint64 {
	return dl.size
}
