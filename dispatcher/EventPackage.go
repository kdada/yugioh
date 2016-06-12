package dispatcher

import "pkg/math/list"

//监听器Id
type ListenerId uint

// Valid 判断监听器id是否有效
func (this ListenerId) Valid() bool {
	return this > 0
}

//事件包
type EventPackage struct {
	Event           IEvent     //事件
	Listeners       list.IList //事件监听器数组,用于存储ListenerNode
	ListenerCounter ListenerId //监听器计数器
}

// NewEventPackage 创建事件包
func NewEventPackage(event IEvent) *EventPackage {
	var eventPackage = new(EventPackage)
	eventPackage.Event = event
	eventPackage.Listeners = list.NewList()
	eventPackage.ListenerCounter = 0
	return eventPackage
}

// ListenerNode 监听器节点
type ListenerNode struct {
	list.LinkedNode             //继承节点
	ListenerId      ListenerId  //监听器id,同一事件包中的监听器id均不同
	Listener        interface{} //监听器
}
