package card

import (
	"yugioh/base"
)

//魔法卡基本数据
type BaseSpellCard struct {
	BaseCard                         //基础信息
	spellCardType base.SpellCardType //魔法卡类型
}

// SpellCardType 返回魔法卡类型
func (this *BaseSpellCard) SpellCardType() base.SpellCardType {
	return this.spellCardType
}
