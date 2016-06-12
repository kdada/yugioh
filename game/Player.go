package game

import (
	"pkg/math/list"
	"yugioh/base"
	"yugioh/card"
)

//卡槽节点
type CardSlotNode struct {
	list.LinkedNode          //链表节点
	CardSlot        CardSlot //卡片槽位
}

//玩家信息
type Player struct {
	PlayerId    int                 //玩家id
	Team        *Team               //玩家队伍id
	Health      int                 //生命值
	GameOver    bool                //是否战败
	MonsterArea [5]*MonsterCardSlot //怪兽区域
	EffectArea  [5]*CardSlot        //效果区域
	HandArea    []card.ICard        //手牌区域
	DeckArea    []card.ICard        //卡组区域
	GraveArea   []card.ICard        //墓地区域
	ExtraArea   []card.ICard        //额外卡组区域
	RemoveArea  []card.ICard        //移除区域
	EffectList  list.IList          //可发动效果列表 EffectNode
}

// NewPlayer 创建玩家信息
func NewPlayer(team *Team, id int) *Player {
	var player = new(Player)
	player.PlayerId = id
	player.Team = team
	player.Health = 8000
	player.GameOver = false
	player.MonsterArea = [5]*MonsterCardSlot{NewMonsterCardSlot(),
		NewMonsterCardSlot(),
		NewMonsterCardSlot(),
		NewMonsterCardSlot(),
		NewMonsterCardSlot()}
	player.EffectArea = [5]*CardSlot{NewCardSlot(base.CardSlotTypeEffect),
		NewCardSlot(base.CardSlotTypeEffect),
		NewCardSlot(base.CardSlotTypeEffect),
		NewCardSlot(base.CardSlotTypeEffect),
		NewCardSlot(base.CardSlotTypeEffect)}
	player.HandArea = make([]card.ICard, 0, 7)
	player.DeckArea = make([]card.ICard, 0, 60)
	player.GraveArea = make([]card.ICard, 0, 20)
	player.ExtraArea = make([]card.ICard, 0, 15)
	player.RemoveArea = make([]card.ICard, 0, 5)
	player.EffectList = list.NewList()
	return player
}

// PopDeckCard 取出卡组顶部的卡牌
func (this *Player) PopDeckCard() (card.ICard,bool) {
	var deckCardCount = len(this.DeckArea)
	if deckCardCount > 0 {
		var card = this.DeckArea[deckCardCount - 1]
		this.DeckArea = this.DeckArea[:deckCardCount-1]
		return card,true
	} 
	return nil,false
}

// AddHandCard 添加卡牌到手牌
func (this *Player) AddHandCard(card card.ICard) {
	this.HandArea = append(this.HandArea,card)
}

// AddValidEffect 添加效果到当前玩家的可发动效果列表
func (this *Player) AddValidEffect(node *EffectNode) {
	this.EffectList.Push(node)
}

// ClearEffect 清空可发动列表
func (this *Player) ClearEffect() {
	this.EffectList.Clear()
}


