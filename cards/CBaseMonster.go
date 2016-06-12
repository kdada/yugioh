package cards

import (
	"pkg/dispatcher"
	"yugioh/base"
	"yugioh/card"
)

//基础怪兽卡片
type CBaseMonster struct {
	card.MonsterCard
	summonNormalListenerId1 dispatcher.ListenerId //主要阶段1普通召唤监听器id
	summonNormalListenerId2 dispatcher.ListenerId //主要阶段2普通召唤监听器id

}

//事件部分

// Init 初始化卡片的自定义的字段
func (this *CBaseMonster) Init(environment card.ICardEnvironment) {

}

// EffectSpeed 返回效果的速度 默认为1
func (this *CBaseMonster) EffectSpeed(environment card.ICardEnvironment, index uint) int {
	return 1
}

// EffectType 返回指定序号的效果的效果类型
func (this *CBaseMonster) EffectType(index uint) base.EffectType {
	switch index {
	case 0:
		return base.EffectTypeSummon
	case 1:
		return base.EffectTypeSummonCover
	default:
		return base.EffectTypeNone
	}
}

// InvokeEffect 执行指定序号的效果
func (this *CBaseMonster) InvokeEffect(environment card.ICardEnvironment, index uint) {
	if this.SceneArea() == base.SceneAreaHand {
		//普通祭品召唤
		var offeringCount = 0
		var star = this.Star()
		switch {
		case star > 4 && star <= 6:
			offeringCount = 1
		case star > 6 && star <= 8:
			offeringCount = 2
		case star > 8:
			offeringCount = 3
		}
		var cards = make([]card.ICard, 0, offeringCount)
		if offeringCount > 0 {
			//选择指定数量的祭品
			cards = append(cards, environment.SelectCards(this.CurrentOwnerId(), environment.OfferingMonsterCards(this.CurrentOwnerId()), offeringCount)...)
		}
		if len(cards) == offeringCount {
			for _, c := range cards {
				//移动到墓地
				environment.MoveCard(c, base.SceneAreaGrave, base.CardStateNone)
			}
			switch index {
			case 0:
				//正面攻击表示
				environment.MoveCard(this, base.SceneAreaMonster, base.CardStateFrontVertical)
			case 1:
				//背面守备表示
				environment.MoveCard(this, base.SceneAreaMonster, base.CardStateBackHorizontal)
			}
		}
	}

	return
}

// OnEnvironmentInited 注册当前卡片需要特殊的事件,在游戏环境初始化完成之后被游戏场景调用
// 游戏开始之前,卡片的特殊效果事件都应该通过此方法被注册到游戏的事件列表中
func (this *CBaseMonster) OnEnvironmentInited(environment card.ICardEnvironment) {
	return
}

// OnMoving 在卡片从一个区域被移动到另一个区域之前,会被游戏场景调用
// 可通过该方法注册一些事件监听器
// 游戏刚开始第一次调用该方法参数必然是from:SceneAreaNone to:SceneAreaDeck
func (this *CBaseMonster) OnMoving(environment card.ICardEnvironment, from base.SceneArea, to base.SceneArea) {
	return
}

// OnMoved 在卡片从一个区域被移动到另一个区域之后,会被场景调用
// 可通过该方法注册一些事件监听器
// 游戏刚开始第一次调用该方法参数必然是from:SceneAreaNone to:SceneAreaDeck
func (this *CBaseMonster) OnMoved(environment card.ICardEnvironment, from base.SceneArea, to base.SceneArea) {
	if to == base.SceneAreaHand {
		//当卡片被加入手牌的时候 注册主要阶段监听器 用于普通召唤
		var listener = func(environment card.ICardEnvironment, from, to base.GameStageType) {
			if environment.CanSummon(base.MonsterSummonTypeNormal) {
				//0号为攻击表示召唤效果,1为里侧守备表示召唤效果
				environment.AddValidEffect(this, 0)
				environment.AddValidEffect(this, 1)
			}
		}
		this.summonNormalListenerId1 = environment.RegisterListener(base.EventNameStageBeforeMainOne, listener)
		this.summonNormalListenerId2 = environment.RegisterListener(base.EventNameStageBeforeMainTwo, listener)
	} else if to != base.SceneAreaHand {
		//当卡片不在手牌中时 移除普通召唤监听器
		environment.RemoveListener(base.EventNameStageBeforeMainOne, this.summonNormalListenerId1)
		environment.RemoveListener(base.EventNameStageBeforeMainTwo, this.summonNormalListenerId2)
	}
	return
}
