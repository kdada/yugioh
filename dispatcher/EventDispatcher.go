package dispatcher

import (
	"pkg/math/list"
	"reflect"
)

//事件分发器
type EventDispatcher struct {
	eventPackages map[string]*EventPackage //事件包map
}

// NewEventDispatcher 创建事件分发器
func NewEventDispatcher() *EventDispatcher {
	var eventDispatcher = new(EventDispatcher)
	eventDispatcher.eventPackages = make(map[string]*EventPackage, 10)
	return eventDispatcher
}

// RegisterEvent 注册事件,所有事件在能够使用之前必须先注册
// 如果已经存在同名的事件,则后注册的事件无效
func (this *EventDispatcher) RegisterEvent(event IEvent) bool {
	if event != nil {
		var _, ok = this.eventPackages[event.Name()]
		if !ok {
			var ep = NewEventPackage(event)
			this.eventPackages[event.Name()] = ep
			return true
		}
	}
	return false
}

// RemoveEvent 移除事件,移除后相应的监听器也会一并移除
// 如果事件存在,则移除并返回true,其他情况返回false
func (this *EventDispatcher) RemoveEvent(eventName string) bool {
	var _, ok = this.eventPackages[eventName]
	if ok {
		delete(this.eventPackages, eventName)
		return true
	}
	return false
}

// DispatchEvent 分发事件
func (this *EventDispatcher) DispatchEvent(eventName string, params ...interface{}) bool {
	var ep, ok = this.eventPackages[eventName]
	if ok {
		ep.Event.OnInvokeListenerStarted(ep.Listeners.Len(), params...)
		var needRemoveNodes = make([]list.ILinkedNode, 0, 10)
		ep.Listeners.Foreach(func(i uint, n list.ILinkedNode) bool {
			var node = n.(*ListenerNode)
			var hold = ep.Event.InvokeListener(node.Listener, node.ListenerId, params...)
			if !hold {
				needRemoveNodes = append(needRemoveNodes, n)
			}
			return true
		})
		//遍历结束后一次性移除所有需要移除的监听器,避免死锁
		for _, node := range needRemoveNodes {
			node.RemoveFromList()
		}
		ep.Event.OnInvokeListenerFinished(params...)
		return true
	}
	return false
}

// RegisterEventListener 注册事件监听器
// 事件必须已经注册,否则无法注册监听器
// 返回事件监听器的id,若id.Valid()为true,则说明注册成功,否则说明注册失败
func (this *EventDispatcher) RegisterEventListener(eventName string, listener interface{}) ListenerId {
	var ep, ok = this.eventPackages[eventName]
	if ok && ep.Event.ValidateListenerType(listener) {
		ep.ListenerCounter++
		var listenerNode = new(ListenerNode)
		listenerNode.Listener = listener
		listenerNode.ListenerId = ep.ListenerCounter
		ep.Listeners.Push(listenerNode)
		return listenerNode.ListenerId
	}
	return 0
}

// RemoveEventListener 移除事件监听器
// 如非必要,请使用RemoveEventListenerById方法
func (this *EventDispatcher) RemoveEventListener(eventName string, listener interface{}) bool {
	var ep, ok = this.eventPackages[eventName]
	if ok && ep.Event.ValidateListenerType(listener) {
		var needRemoveNode list.ILinkedNode = nil
		ep.Listeners.Foreach(func(i uint, n list.ILinkedNode) bool {
			var node = n.(*ListenerNode)
			if reflect.ValueOf(node.Listener).Pointer() == reflect.ValueOf(listener).Pointer() {
				needRemoveNode = n
				return false
			}
			return true
		})
		if needRemoveNode != nil {
			return needRemoveNode.RemoveFromList()
		}
	}
	return false
}

// RemoveEventListenerById 使用监听器id移除监听器
func (this *EventDispatcher) RemoveEventListenerById(eventName string, listenerId ListenerId) bool {
	var ep, ok = this.eventPackages[eventName]
	if ok {
		var needRemoveNode list.ILinkedNode = nil
		ep.Listeners.Foreach(func(i uint, n list.ILinkedNode) bool {
			var node = n.(*ListenerNode)
			if node.ListenerId == listenerId {
				needRemoveNode = n
				return false
			}
			return true
		})
		if needRemoveNode != nil {
			return needRemoveNode.RemoveFromList()
		}
	}
	return false
}
