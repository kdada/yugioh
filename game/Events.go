package game

import (
	"pkg/dispatcher"
	"yugioh/base"
	"yugioh/card"
)

////////////////////////////阶段切换事件部分////////////////////////////开始
//事件监听器
//environment:运行环境
//from:原阶段
//to:当前阶段
//return:返回是否保留当前监听器,返回true则保留,否则执行完成之后就移除当前监听器
//type StageChangedListener func(environment card.ICardEnvironment,from ,to base.GameStageType) bool

//阶段切换事件
type StageChangedEvent struct {
	dispatcher.BaseEvent
}

// NewStageChangedEvent 创建基本阶段变更事件
func NewStageChangedEvent(name base.EventName) *StageChangedEvent {
	var event = new(StageChangedEvent)
	event.Init(string(name))
	return event
}

//自动阶段切换事件
//自动阶段切换事件结束后,会自动使用environment进入下一个阶段
type StageAutoChangedEvent struct {
	StageChangedEvent
}

// StageAutoChangedEvent 创建自动阶段变更事件
func NewStageAutoChangedEvent(name base.EventName) *StageAutoChangedEvent {
	var event = new(StageAutoChangedEvent)
	event.Init(string(name))
	return event
}

// OnInvokeListenerFinished 在所有监听器执行完毕后被调用
func (this *StageChangedEvent) OnInvokeListenerFinished(params ...interface{}) {
	var environment, ok = params[0].(card.ICardEnvironment)
	if ok {
		if environment.Stage() == base.GameStageTypeEnd {
			//如果当前为最后一个阶段则直接切换到下一回合
			environment.NextRange()
		} else {
			environment.NextStage()
		}
	}
}

////////////////////////////阶段切换事件部分////////////////////////////结束
