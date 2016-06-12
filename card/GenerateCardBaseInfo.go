package card

import (
	"yugioh/base"
)

//设置基础信息接口
type setBaseCardInfo interface {
	setBaseCardInfoHiddenField(base IBaseCard)
}

//设置游戏卡片信息接口
type setGameCardInfo interface {
	setGameCardInfoHiddenField(origin IBaseCard, cardId base.CardId, ownerId int)
}

// SetCardBaseInfo 设置游戏卡片信息
// card必须继承自MonsterCard或SpellCard或TrapCard
// baseCard必须是BaseMonsterCard或BaseSepllCard或BaseTrapCard
// cardId为游戏时卡片id
// ownerId为拥有者id
func SetCardBaseInfo(card ICard, baseCard IBaseCard, cardId base.CardId, ownerId int) bool {
	if card != nil && baseCard != nil {
		baseInfo, ok := card.(setBaseCardInfo)
		gameInfo, sok := card.(setGameCardInfo)
		if ok && sok {
			baseInfo.setBaseCardInfoHiddenField(baseCard)
			gameInfo.setGameCardInfoHiddenField(baseCard, cardId, ownerId)
			return true
		}
	}
	return false
}

// NewBaseMonsterCard 创建基础怪兽信息
func NewBaseMonsterCard(cardNo base.CardNo,
	cardName string,
	cardType base.CardType,
	effectCount uint8,
	monsterProperty base.MonsterProperty,
	monsterRace base.MonsterRace,
	monsterSummonType base.MonsterSummonType,
	monsterAttackType base.MonsterAttackType,
	monsterDefenceType base.MonsterDefenceType,
	star,
	attack,
	defence int) *BaseMonsterCard {
	var card = new(BaseMonsterCard)
	card.cardNo = cardNo
	card.cardName = cardName
	card.cardType = cardType
	card.effectCount = effectCount
	card.monsterProperty = monsterProperty
	card.monsterRace = monsterRace
	card.monsterSummonType = monsterSummonType
	card.monsterAttackType = monsterAttackType
	card.monsterDefenceType = monsterDefenceType
	card.star = star
	card.attack = attack
	card.defence = defence
	return card
}

// setBaseCardInfoHiddenField 设置基础信息
func (this *BaseMonsterCard) setBaseCardInfoHiddenField(b IBaseCard) {
	var info, ok = b.(*BaseMonsterCard)
	if ok {
		info.cardType = base.CardTypeMonster

		this.cardNo = info.cardNo
		this.cardName = info.cardName
		this.cardType = info.cardType
		this.effectCount = info.effectCount
		this.monsterProperty = info.monsterProperty
		this.monsterRace = info.monsterRace
		this.monsterSummonType = info.monsterSummonType
		this.monsterAttackType = info.monsterAttackType
		this.monsterDefenceType = info.monsterDefenceType
		this.star = info.star
		this.attack = info.attack
		this.defence = info.defence
	}
}

// NewBaseSpellCard 创建基础魔法卡信息
func NewBaseSpellCard(cardNo base.CardNo,
	cardName string,
	cardType base.CardType,
	effectCount uint8,
	spellCardType base.SpellCardType) *BaseSpellCard {
	var card = new(BaseSpellCard)
	card.cardNo = cardNo
	card.cardName = cardName
	card.cardType = cardType
	card.effectCount = effectCount
	card.spellCardType = spellCardType
	return card
}

// setBaseCardInfoHiddenField 设置基础信息
func (this *BaseSpellCard) setBaseCardInfoHiddenField(b IBaseCard) {
	var info, ok = b.(*BaseSpellCard)
	if ok {
		info.cardType = base.CardTypeSpell

		this.cardNo = info.cardNo
		this.cardName = info.cardName
		this.cardType = info.cardType
		this.effectCount = info.effectCount
		this.spellCardType = info.spellCardType
	}
}

// NewBaseTrapCard 创建基础魔法卡信息
func NewBaseTrapCard(cardNo base.CardNo,
	cardName string,
	cardType base.CardType,
	effectCount uint8,
	trapCardType base.TrapCardType) *BaseTrapCard {
	var card = new(BaseTrapCard)
	card.cardNo = cardNo
	card.cardName = cardName
	card.cardType = cardType
	card.effectCount = effectCount
	card.trapCardType = trapCardType
	return card
}

// setBaseCardInfoHiddenField 设置基础信息
func (this *BaseTrapCard) setBaseCardInfoHiddenField(b IBaseCard) {
	var info, ok = b.(*BaseTrapCard)
	if ok {
		info.cardType = base.CardTypeTrap

		this.cardNo = info.cardNo
		this.cardName = info.cardName
		this.cardType = info.cardType
		this.effectCount = info.effectCount
		this.trapCardType = info.trapCardType
	}
}

// 设置游戏时卡片信息
func (this *GameCard) setGameCardInfoHiddenField(origin IBaseCard, cardId base.CardId, ownerId int) {
	this.origin = origin
	this.cardId = cardId
	this.sceneArea = base.SceneAreaNone
	this.cardState = base.CardStateNone
	this.ownerId = ownerId
	this.currentOwnerId = ownerId
	this.relatedCards = make([]IGameCard, 0)
}
