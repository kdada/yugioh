package card

import (
	"yugioh/base"
)

//卡片基本数据
//实现:IBaseCard
type BaseCard struct {
	cardNo      base.CardNo   //卡片编号,同样的卡编号相同
	cardName    string        //卡片名称
	cardType    base.CardType //卡片类型
	effectCount uint8         //当前卡片拥有的效果数量
	//CardDesc string
	//CardImg  string
}

// CardNo 返回当前卡片的编号
func (this *BaseCard) CardNo() base.CardNo {
	return this.cardNo
}

// CardName 返回当前卡片名称
func (this *BaseCard) CardName() string {
	return this.cardName
}

// CardType 返回当前卡片的类型:怪兽,魔法,陷阱
func (this *BaseCard) CardType() base.CardType {
	return this.cardType
}

// EffectCount 返回当前卡片拥有的效果数量
func (this *BaseCard) EffectCount() uint {
	return uint(this.effectCount)
}

// EffectType 返回指定序号的效果的效果类型
func (this *BaseCard) EffectType(index uint) base.EffectType {
	return base.EffectTypeNone
}