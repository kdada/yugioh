package dispatcher

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

//测试事件分配器功能
func TestEventDispatcher(t *testing.T) {
	var ta = (*TestAlert)(t)
	var dispatcher = NewEventDispatcher()
	var event = NewBaseEvent("EventXXX")

	dispatcher.RegisterEvent(event)
	ta.Assert(len(dispatcher.eventPackages) != 1, "1.注册事件失败")

	var listener = func(name string) bool {
		fmt.Println("aa:", name)
		return true
	}
	var listenerId = dispatcher.RegisterEventListener(event.Name(), listener)
	ta.Assert(!listenerId.Valid(), "1.监听器注册失败")

	var listenerId2 = dispatcher.RegisterEventListener(event.Name(), func(name string) {
		fmt.Println("bb:", name)
	})
	ta.Assert(!listenerId2.Valid(), "2.监听器注册失败")

	var result = dispatcher.DispatchEvent(event.Name(), event.Name())
	ta.Assert(!result, "1.分发事件失败")

	var ok = dispatcher.RemoveEventListener(event.Name(), listener)
	ta.Assert(!ok, "1.监听器移除失败")

	result = dispatcher.DispatchEvent(event.Name(), event.Name())
	ta.Assert(!result, "2.分发事件失败")

	listenerId = dispatcher.RegisterEventListener(event.Name(), listener)
	ta.Assert(!listenerId.Valid(), "3.监听器注册失败")

	result = dispatcher.DispatchEvent(event.Name(), event.Name())
	ta.Assert(!result, "3.分发事件失败")

	ok = dispatcher.RemoveEventListenerById(event.Name(), listenerId2)
	ta.Assert(!ok, "2.监听器移除失败")

	result = dispatcher.DispatchEvent(event.Name(), event.Name())
	ta.Assert(!result, "4.分发事件失败")

	ok = dispatcher.RemoveEvent(event.Name())
	ta.Assert(!ok, "1.事件移除失败")

	result = dispatcher.DispatchEvent(event.Name(), event.Name())
	ta.Assert(result, "5.分发事件失败")
	//Output:
	//aa: Event
	//bb: Event
	//bb: Event
	//bb: Event
	//aa: Event
	//aa: Event
}

//监听器移除情况
func ExampleEventListener() {
	var dispatcher = NewEventDispatcher()
	var event = NewBaseEvent("Event")
	dispatcher.RegisterEvent(event)
	dispatcher.RegisterEventListener(event.Name(), (func(name string) bool {
		fmt.Println("xxxx removed")
		return false
	}))
	dispatcher.RegisterEventListener(event.Name(), (func(name string) bool {
		fmt.Println("xxxx2")
		return true
	}))

	dispatcher.RegisterEventListener(event.Name(), (func(name string) bool {
		fmt.Println("xxxx3")
		return true
	}))
	dispatcher.DispatchEvent(event.Name(), event.Name())

	dispatcher.DispatchEvent(event.Name(), event.Name())
	//Output:
	//xxxx removed
	//xxxx2
	//xxxx3
	//xxxx2
	//xxxx3
}
