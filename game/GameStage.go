package game

import (
	"yugioh/base"
)

//游戏阶段节点
type GameStage struct {
	GameStageType base.GameStageType //游戏阶段
	GameStageName base.EventName     //游戏阶段事件名称
	Skip          bool               //跳过当前阶段
}

// NewGameStage 创建游戏阶段
func NewGameStage(stageType base.GameStageType, eventName base.EventName) *GameStage {
	var stage = new(GameStage)
	stage.GameStageType = stageType
	stage.GameStageName = eventName
	stage.Skip = false
	return stage
}

// Resume 恢复阶段
func (this *GameStage) Resume() {
	this.Skip = false
}
