package card

import (
	"pkg/dispatcher"
	"yugioh/base"
)

//所有效果的实现都必须通过本接口的实例进行,不应该直接修改卡片数据
type ICardEnvironment interface {
	// Time 返回当前回合数
	Time() int
	// Stage 返回当前回合所处的阶段
	Stage() base.GameStageType
	// CurrentPlayerId 返回当前回合的玩家id
	CurrentPlayerId() int
	// NextStage 切换到下一阶段没有被跳过的阶段
	NextStage()
	// GoToStage 切换到指定的阶段
	GoToStage(stageType base.GameStageType)
	// NextRange 进入下一回合,如果当前处于第0回合,则自动初始化回合信息
	NextRange()
	// DrawCard 指定player抽count张卡
	DrawCard(playerId, count int)
	// AddHealth 增加玩家生命值,count为负数则为减
	AddHealth(playerId, count int)
	// AddValidEffect 给当前卡片所属玩家添加可发动效果
	AddValidEffect(card ICard, effectIndex uint)
	// InvokeEffectIndex 将效果加入连锁列表,等待发动卡片指定效果
	AddEffectToChain(card ICard, effectIndex uint)
	// InvokeEffect 发动效果
	InvokeEffect(card ICard, effectIndex uint)
	// TestWin 检查是否存在胜利队伍
	TestWin()
	// CanSummon 检查场景是否允许进行指定类型的召唤
	CanSummon(summonType base.MonsterSummonType) bool
	// MoveCard 移动卡片位置,只有to为怪兽或魔法陷阱区域state才有效
	MoveCard(card ICard, to base.SceneArea, state base.CardState)
	// OfferingMonsterCards 返回指定玩家可以作为祭品的怪兽卡片
	OfferingMonsterCards(playerId int) []ICard
	// SelectCards 让玩家选择从卡片中选择指定数量的卡片
	SelectCards(playerId int, cards []ICard, count int) []ICard
	// DisptachEvent 分发事件
	DisptachEvent(eventName base.EventName, params ...interface{})
	// RegisterEvent 注册事件
	RegisterEvent(event dispatcher.IEvent)
	// RegisterListener 注册事件监听器
	RegisterListener(eventName base.EventName, listener interface{}) dispatcher.ListenerId
	// RemoveListener 移除事件监听器
	RemoveListener(eventName base.EventName, listenerId dispatcher.ListenerId)
}
