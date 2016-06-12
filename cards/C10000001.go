package cards

import (
	"yugioh/base"
	"yugioh/card"
)

//卡片Xxxx
type C10000001 struct {
	CBaseMonster
	a int
}

//注册卡片信息到卡表
var _ = func() int {
	var baseInfo = new(cardBaseInfo)
	baseInfo.cardNo = 10000001
	baseInfo.cardName = "没名字"
	baseInfo.cardType = base.CardTypeMonster
	baseInfo.effectCount = 1
	baseInfo.monsterProperty = base.MonsterPropertyDivine
	baseInfo.monsterRace = base.MonsterRaceWyrm
	baseInfo.monsterSummonType = base.MonsterSummonTypeNormal | base.MonsterSummonTypeEffect
	baseInfo.monsterAttackType = base.MonsterAttackTypeNormal | base.MonsterAttackTypeDirect
	baseInfo.monsterDefenceType = base.MonsterDefenceTypeNormal
	baseInfo.star = 1
	baseInfo.attack = 4000
	baseInfo.defence = 4000
	registerCard(baseInfo, new(C10000001))
	return 0
}()

//事件部分

// Init 初始化卡片的自定义的字段
func (this *C10000001) Init(environment card.ICardEnvironment) {
	this.a = 100
}

// EffectSpeed 返回效果的速度 默认为1
func (this *C10000001) EffectSpeed(environment card.ICardEnvironment, index uint) int {
	return 1
}

// InvokeEffect 执行指定序号的效果
func (this *C10000001) InvokeEffect(environment card.ICardEnvironment, index uint) {
	return
}

// OnEnvironmentInited 注册当前卡片需要特殊的事件,在游戏环境初始化完成之后被游戏场景调用
// 游戏开始之前,卡片的特殊效果事件都应该通过此方法被注册到游戏的事件列表中
func (this *C10000001) OnEnvironmentInited(environment card.ICardEnvironment) {
	return
}

// OnMoving 在卡片从一个区域被移动到另一个区域之前,会被游戏场景调用
// 可通过该方法注册一些事件监听器
// 游戏刚开始第一次调用该方法参数必然是from:SceneAreaNone to:SceneAreaDeck
func (this *C10000001) OnMoving(environment card.ICardEnvironment, from base.SceneArea, to base.SceneArea) {
	return
}

// OnMoved 在卡片从一个区域被移动到另一个区域之后,会被场景调用
// 可通过该方法注册一些事件监听器
// 游戏刚开始第一次调用该方法参数必然是from:SceneAreaNone to:SceneAreaDeck
func (this *C10000001) OnMoved(environment card.ICardEnvironment, from base.SceneArea, to base.SceneArea) {
	return
}
