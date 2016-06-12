package game

import (
	"yugioh/base"
	"yugioh/card"
)

//卡片槽位
type CardSlot struct {
	CardSlotType base.CardSlotType //卡槽类型
	Valid        bool              //卡片槽位是否有效
	Card         card.ICard        //卡片

}

// NewCardSlot 创建卡槽
func NewCardSlot(slotType base.CardSlotType) *CardSlot {
	var slot = new(CardSlot)
	slot.CardSlotType = slotType
	slot.Valid = true
	return slot
}

// CanSetCard 返回当前槽位能否设置卡片
func (this *CardSlot) CanSetCard() bool {
	return this.Card == nil && this.Valid
}

//怪兽卡片槽位
type MonsterCardSlot struct {
	CardSlot
	AttackCount    int                //可攻击次数
	AttackDirectly bool               //能否直接攻击
	AttackSlots    []*MonsterCardSlot //可攻击的对象的槽位

}

// MonsterCardSlot 创建怪兽卡槽
func NewMonsterCardSlot() *MonsterCardSlot {
	var slot = new(MonsterCardSlot)
	slot.CardSlotType = base.CardSlotTypeMonster
	slot.Valid = true
	slot.AttackSlots = make([]*MonsterCardSlot, 0, 5)
	return slot
}

// RemoveAllAttackSlots 移除所有可攻击对象槽位
func (this *MonsterCardSlot) RemoveAllAttackSlots() {
	this.AttackSlots = this.AttackSlots[:0]
}

// AddAttackSlot 添加可攻击槽位
func (this *MonsterCardSlot) AddAttackSlot(slot *MonsterCardSlot) {
	this.AttackSlots = append(this.AttackSlots, slot)
}

// AttackSlotCount 返回可攻击槽位的数量
func (this *MonsterCardSlot) AttackSlotCount() int {
	return len(this.AttackSlots)
}
