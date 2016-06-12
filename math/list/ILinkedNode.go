package list

//链表节点
type ILinkedNode interface {
	// PreNode 返回前一个节点
	PreNode() ILinkedNode
	// setPreNode 设置前一个节点
	setPreNode(node ILinkedNode)
	// NextNode 返回后一个节点
	NextNode() ILinkedNode
	// setNextNode 设置后一个节点
	setNextNode(node ILinkedNode)
	// setList 设置节点所属的链表
	setList(list IList)
	// InList 判断当前节点是否在链表中
	InList() bool
	// RemoveFromList 从链表中移除当前节点
	// 注意:该方法会使用Node所属的list的移除方法,在遍历list时使用会产生死锁
	RemoveFromList() bool
}
