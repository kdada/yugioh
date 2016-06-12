package list

//链表节点
//实现:ILinkNode
type LinkedNode struct {
	preNode  ILinkedNode //前一个节点
	nextNode ILinkedNode //下一个节点
	list     IList       //节点所属的list
}

// PreNode 前一个节点
func (this *LinkedNode) PreNode() ILinkedNode {
	return this.preNode
}

// setPreNode 设置前一个节点
func (this *LinkedNode) setPreNode(node ILinkedNode) {
	this.preNode = node
}

// NextNode 后一个节点
func (this *LinkedNode) NextNode() ILinkedNode {
	return this.nextNode
}

// setNextNode 设置后一个节点
func (this *LinkedNode) setNextNode(node ILinkedNode) {
	this.nextNode = node
}

// setList 设置节点所属的链表
func (this *LinkedNode) setList(list IList) {
	this.list = list
}

// InList 当前节点是否在链表中
func (this *LinkedNode) InList() bool {
	return this.list != nil
}

// RemoveFromList 从链表中移除当前节点
func (this *LinkedNode) RemoveFromList() bool {
	if this.InList() {
		return this.list.removeNode(this)
	}
	return false
}
