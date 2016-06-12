// list 包实现了一个多线程安全的链表,并且支持作为栈或队列使用
package list

import "sync"

//链表,多线程安全
//实现:IList
type List struct {
	length    uint          //链表长度
	startNode ILinkedNode   //链表起始节点,构成循环链表
	rwMutex   *sync.RWMutex //读写锁
}

// NewList 创建一个链表
func NewList() *List {
	var list = new(List)
	list.startNode = new(LinkedNode)
	list.startNode.setNextNode(list.startNode)
	list.startNode.setPreNode(list.startNode)
	list.startNode.setList(list)
	list.rwMutex = new(sync.RWMutex)
	return list
}

// Len 返回链表长度
func (this *List) Len() uint {
	return this.length
}

// addNode 添加节点,List中所有添加操作都必须调用本方法
// node将添加在preNode之后
func (this *List) addNode(preNode, node ILinkedNode) bool {
	//写锁定
	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()
	if preNode.InList() && !node.InList() {
		preNode.NextNode().setPreNode(node)
		node.setNextNode(preNode.NextNode())
		preNode.setNextNode(node)
		node.setPreNode(preNode)
		node.setList(this)
		//增加链表长度
		this.length++
		return true
	}
	return false
}

// removeNode 删除节点,List中所有删除操作都必须调用本方法
func (this *List) removeNode(node ILinkedNode) bool {
	//写锁定
	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()
	if node.InList() {
		node.PreNode().setNextNode(node.NextNode())
		node.NextNode().setPreNode(node.PreNode())
		node.setNextNode(nil)
		node.setPreNode(nil)
		node.setList(nil)
		//减少链表长度
		this.length--
		return true
	}
	return false
}
// Clear 清空所有节点
func (this *List) Clear() {
	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()
	this.length = 0
	this.startNode.setNextNode(this.startNode)
	this.startNode.setPreNode(this.startNode)
}


// Push 在尾部添加节点,不能添加nil节点
func (this *List) Push(node ILinkedNode) bool {
	if node != nil {
		var lastNode = this.startNode.PreNode()
		return this.addNode(lastNode, node)
	}
	return false
}

// Pop 从尾部删除节点
func (this *List) Pop() (ILinkedNode, bool) {
	if this.length > 0 {
		var lastNode = this.startNode.PreNode()
		this.removeNode(lastNode)
		return lastNode, true
	}
	return nil, false
}

// Enqueue 在尾部添加节点,不能添加nil节点
func (this *List) Enqueue(node ILinkedNode) bool {
	return this.Push(node)
}

// Dequeue 从首部删除节点
func (this *List) Dequeue() (ILinkedNode, bool) {
	if this.length > 0 {
		var firstNode = this.startNode.NextNode()
		this.removeNode(firstNode)
		return firstNode, true
	}
	return nil, false
}

// Insert 在指定位置添加节点,不能添加nil节点
func (this *List) Insert(index uint, node ILinkedNode) bool {
	if node != nil {
		if index == this.length {
			return this.Push(node)
		} else if index < this.length {
			var preNode ILinkedNode = nil
			this.Foreach(func(i uint, n ILinkedNode) bool {
				if i == index {
					preNode = n.PreNode()
					return false
				}
				return true
			})
			return this.addNode(preNode, node)
		}
	}
	return false
}

// Remove 删除指定位置上的节点
func (this *List) Remove(index uint) (ILinkedNode, bool) {
	if index >= this.length {
		return nil, false
	}
	var removeNode ILinkedNode = nil
	this.Foreach(func(i uint, node ILinkedNode) bool {
		if i == index {
			removeNode = node
			return false
		}
		return true
	})
	return removeNode, this.removeNode(removeNode)
}

// Index 获取链表中的一项
func (this *List) Index(index uint) (ILinkedNode, bool) {
	//读锁定
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	if index >= this.length {
		return nil, false
	}
	var result ILinkedNode = nil
	this.Foreach(func(i uint, node ILinkedNode) bool {
		if i == index {
			result = node
			return false
		}
		return true
	})
	return result, true
}

// Foreach 遍历链表中的每个节点,each方法的返回值确定是否继续遍历,返回true则继续遍历
// 遍历过程中,不能添加或移除任何节点,否则会导致死锁
func (this *List) Foreach(each func(i uint, n ILinkedNode) bool) {
	//读锁定
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	if this.length > 0 {
		var currentNode = this.startNode.NextNode()
		var currentIndex uint = 0
		for (currentNode != nil) && (currentNode != this.startNode) {
			var goOn = each(currentIndex, currentNode)
			if !goOn {
				break
			}
			currentNode = currentNode.NextNode()
			currentIndex++
		}
	}

}
