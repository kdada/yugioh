package dispatcher

//事件接口
type IEvent interface {
	// Name 返回事件名称
	Name() string
	// ValidateListenerType 验证事件监听器是否与当前事件相匹配
	ValidateListenerType(listener interface{}) bool
	// InvokeListener 执行监听器,并确定该事件监听器是否应该被保留,返回false则被移除
	InvokeListener(listener interface{}, id ListenerId, params ...interface{}) bool
	// OnInvokeListenerStarted 在执行所有监听器之前被调用
	OnInvokeListenerStarted(listenersCount uint, params ...interface{})
	// OnInvokeListenerFinished 在所有监听器执行完毕后被调用
	OnInvokeListenerFinished(params ...interface{})
}
