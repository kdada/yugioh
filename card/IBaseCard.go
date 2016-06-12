package card

import (
	"yugioh/base"
)

//卡片接口,所有卡牌均实现该接口
type IBaseCard interface {
	// CardNo 返回当前卡片的编号
	CardNo() base.CardNo
	// CardName 返回当前卡片名称
	CardName() string
	// CardType 返回当前卡片的类型:怪兽,魔法,陷阱
	CardType() base.CardType
	// EffectCount 返回当前卡片拥有的效果数量
	EffectCount() uint
	// EffectType 返回指定序号的效果的效果类型
	EffectType(index uint) base.EffectType
}
