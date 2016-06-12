package card

import (
	"yugioh/base"
)

//陷阱卡基本信息接口
type ITrapCard interface {
	IBaseCard
	// TrapCardType 返回陷阱卡类型
	TrapCardType() base.TrapCardType
}
