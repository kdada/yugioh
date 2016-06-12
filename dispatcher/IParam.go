package dispatcher

//事件回调参数接口
//params会被传递给每一个监听器,当params被一个监听器修改后,后续的监听器会使用修改后的值
//type IParam interface {
//	// EventName 返回当前事件名称
//	EventName() string
//	// EventDispatcher 返回当前事件所属的事件分发器
//	EventDispatcher() IEventDispatcher
//	// NeedRemove 确定当前监听器是否需要移除
//	NeedRemove() bool
//	// SetNeedRemove 设置是否需要移除当前监听器
//	SetNeedRemove(remove bool)
//	// Resume 恢复参数数据,每次被传递给一个监听器之前,都会调用该方法对参数进行重置
//	// 默认情况下,Resume方法会调用SetNeedRemove(false)以避免影响之后的监听器
//	// 因此在Resume中设置参数的值
//	Resume()
//}
