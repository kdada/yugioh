package dispatcher

////基础参数
//type BaseParam struct {
//	name       string           //事件名称
//	dispatcher IEventDispatcher //事件分发器
//	remove     bool             //是否移除当前事件回调
//}

//// NewBaseParam 创建基础参数
//func NewBaseParam(eventName string, eventDispatcher IEventDispatcher) *BaseParam {
//	var param = new(BaseParam)
//	param.Init(eventName, eventDispatcher)
//	return param
//}

//// Init 基础参数初始化,默认情况下NeedRemove()为false
//// eventName:事件名称
//// eventDispatcher:事件所属的分发器
//func (this *BaseParam) Init(eventName string, eventDispatcher IEventDispatcher) {
//	this.name = eventName
//	this.dispatcher = eventDispatcher
//	this.remove = false
//}

//// EventName 返回当前事件名称
//func (this *BaseParam) EventName() string {
//	return this.name
//}

//// EventDispatcher 当前事件所属的事件分发器
//func (this *BaseParam) EventDispatcher() IEventDispatcher {
//	return this.dispatcher
//}

//// NeedRemove 当前监听器是否需要移除
//func (this *BaseParam) NeedRemove() bool {
//	return this.remove
//}

//// SetNeedRemove 设置是否需要移除当前监听器
//func (this *BaseParam) SetNeedRemove(remove bool) {
//	this.remove = remove
//}

//// Resume 恢复参数数据,在参数被传递给一个监听器之前，都会调用该方法对参数进行重置
//// 默认情况下,Resume方法会调用SetNeedRemove(false)以避免影响之后的监听器
//func (this *BaseParam) Resume() {
//	this.SetNeedRemove(false)
//}
