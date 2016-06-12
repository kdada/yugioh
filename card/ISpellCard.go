package card

import (
	"yugioh/base"
)

//魔法卡基本信息接口
type ISpellCard interface {
	IBaseCard
	// SpellCardType 返回魔法卡类型
	SpellCardType() base.SpellCardType
}
