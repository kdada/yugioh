package list

//链表接口
type IList interface {
	// addNode 将node添加到preNode之后
	addNode(preNode,node ILinkedNode) bool
	// removeNode 移除node
	removeNode(node ILinkedNode) bool
	// Clear 清空所有节点
	Clear()
	// Len 返回链表长度
	Len() uint
	// Push 在尾部添加节点,不能添加nil节点
	Push(node ILinkedNode) bool
	// Pop 从尾部删除节点
	Pop() (ILinkedNode, bool)
	// Enqueue 在尾部添加节点,不能添加nil节点
	Enqueue(node ILinkedNode) bool
	// Dequeue 从链表首部删除节点,并返回该节点
	Dequeue() (ILinkedNode, bool)
	// Insert 在指定index位置添加node,不能添加nil节点
	Insert(index uint, node ILinkedNode) bool
	// Remove 删除指定index位置上的节点
	Remove(index uint) (ILinkedNode, bool)
	// Index 获取链表中指定index的节点
	Index(index uint) (ILinkedNode, bool)
	// Foreach 遍历链表中的每个节点,each方法的返回值确定是否继续遍历,返回true则继续遍历
	// 遍历过程中,不能添加或移除任何节点,否则会造成死锁等情况
	Foreach(each func(index uint, node ILinkedNode) bool)
}