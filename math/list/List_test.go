package list

import (
	"fmt"
	"testing"
)

type TestAlert testing.T

func (this *TestAlert) Assert(needAlert bool, info string) {
	if needAlert {
		this.Error(info)
	}
}

type TestNode struct {
	LinkedNode
	data uint
}

//测试链表  作为栈使用
func TestListStack(t *testing.T) {
	var ta = (*TestAlert)(t)
	var nodes = [100]*TestNode{}
	var list = NewList()
	var i uint = 0
	for ; i < 100; i++ {
		var node = new(TestNode)
		node.data = i
		nodes[i] = node
		ta.Assert(node.InList(), "1.节点错误")
		var ok = list.Push(node)
		ta.Assert(!ok, "链表压入错误")
		ta.Assert(!node.InList(), "2.节点错误")
		ta.Assert(list.length != i+1, "1.链表长度错误")
	}
	ta.Log("链表长度:", list.Len())
	i = 100
	for ; i > 0; i-- {
		var node, ok = list.Pop()
		ta.Assert(!ok, "链表弹出错误")
		ta.Assert(node.InList(), "3.节点错误")
		ta.Assert(list.length != i-1, "2.链表长度错误")
		s, ok := node.(*TestNode)
		ta.Assert(s.data != i-1, "节点数值错误")
	}
}

//测试链表  作为队列使用
func TestListQueue(t *testing.T) {
	var ta = (*TestAlert)(t)
	var nodes = [100]*TestNode{}
	var list = NewList()
	var i uint = 0
	for ; i < 100; i++ {
		var node = new(TestNode)
		node.data = i
		nodes[i] = node
		ta.Assert(node.InList(), "1.节点错误")
		var ok = list.Enqueue(node)
		ta.Assert(!ok, "链表压入错误")
		ta.Assert(!node.InList(), "2.节点错误")
		ta.Assert(list.length != i+1, "1.链表长度错误")
	}
	ta.Log("链表长度:", list.Len())
	i = 100
	for ; i > 0; i-- {
		var node, ok = list.Dequeue()
		ta.Assert(!ok, "链表弹出错误")
		ta.Assert(node.InList(), "3.节点错误")
		ta.Assert(list.length != i-1, "2.链表长度错误")
		s, ok := node.(*TestNode)
		ta.Assert(s.data != 100-i, "节点数值错误")
	}
}

//测试链表
func TestList(t *testing.T) {
	var ta = (*TestAlert)(t)

	var nodes = [100]*TestNode{}
	var list = NewList()
	var i uint = 0
	for ; i < 100; i++ {
		var node = new(TestNode)
		node.data = i
		nodes[i] = node
	}

	var ok = list.Insert(0, nodes[0])
	ta.Assert(!ok, "1.插入出错")
	ok = list.Insert(1, nodes[0])
	ta.Assert(ok, "2.插入出错")
	ok = list.Insert(0, nodes[1])
	ta.Assert(!ok, "3.插入出错")
	ok = list.Insert(1, nodes[2])
	ta.Assert(!ok, "4.插入出错")

	node, ok := list.Remove(1)
	ta.Assert((!ok) || (list.Len() != 2), "1.删除出错")
	var tn = node.(*TestNode)
	ta.Assert(tn.data != 2, "2.删除出错")
	nodes[1].RemoveFromList()
	ta.Assert((list.Len() != 1), "3.删除出错")
	node, ok = list.Index(0)
	n, ok := node.(*TestNode)
	ta.Assert(n.data != 0, "4.删除错误")

}

func ExampleForeach() {
	var nodes = [100]*TestNode{}
	var list = NewList()
	var i uint = 0
	for ; i < 100; i++ {
		var node = new(TestNode)
		node.data = i
		nodes[i] = node
	}
	list.Insert(0, nodes[0])
	list.Insert(0, nodes[1])
	list.Insert(1, nodes[2])
	var node, _ = list.Index(0)
	fmt.Println(0, node.(*TestNode).data)

	list.Foreach(func(index uint, node ILinkedNode) bool {
		fmt.Println(index, node.(*TestNode).data)
		return true
	})
	
	list.Clear()
	
	list.Foreach(func(index uint, node ILinkedNode) bool {
		fmt.Println(index, node.(*TestNode).data)
		return true
	})
	//Output:
	//0 1
	//0 1
	//1 2
	//2 0
}
