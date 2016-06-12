package dispatcher

import (
	"reflect"
)

//基本事件监听器,监听器应该返回一个bool值,也可以不返回任何内容,不返回则默认为返回true
//返回是否保留当前监听器,返回值为false移除当前监听器
//type BaseEventListener func() bool

//事件结构
type BaseEvent struct {
	name string //事件名称
}

// NewBaseEvent 创建基础事件
func NewBaseEvent(name string) *BaseEvent {
	var event = new(BaseEvent)
	event.name = name
	return event
}

//初始化事件
func (this *BaseEvent) Init(name string) {
	this.name = name
}

// Name 返回事件名称
func (this *BaseEvent) Name() string {
	return this.name
}

// ValidateListenerType 验证事件监听器是否与当前事件相匹配
// 此处并不做验证,如果需要验证则需要继承并重写本方法
func (this *BaseEvent) ValidateListenerType(listener interface{}) bool {
	return true
}

// InvokeListener 执行事件监听器
func (this *BaseEvent) InvokeListener(listener interface{}, id ListenerId, params ...interface{}) bool {
	var listenerType = reflect.ValueOf(listener)
	if listenerType.Kind() == reflect.Func && (listenerType.Type().NumIn() == len(params)) {
		var valueParams = make([]reflect.Value, len(params))
		for i := 0; i < len(valueParams); i++ {
			valueParams[i] = reflect.ValueOf(params[i])
		}
		var result = listenerType.Call(valueParams)
		if len(result) == 1 {
			var v = result[0]
			if v.Kind() == reflect.Bool {
				return v.Bool()
			}
		}
		return true
	}
	//调用出现错误则返回false,移除该监听器
	return false
}

// OnInvokeListenerStarted 在执行所有监听器之前被调用
func (this *BaseEvent) OnInvokeListenerStarted(listenersCount uint, params ...interface{}) {

}

// OnInvokeListenerFinished 在所有监听器执行完毕后被调用
func (this *BaseEvent) OnInvokeListenerFinished(params ...interface{}) {

}
