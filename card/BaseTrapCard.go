package card

import (
	"yugioh/base"
)

//陷阱卡基本数据
type BaseTrapCard struct {
	BaseCard                       //基础信息
	trapCardType base.TrapCardType //陷阱卡类型
}

// TrapCardType 返回陷阱卡类型
func (this *BaseTrapCard) TrapCardType() base.TrapCardType {
	return this.trapCardType
}
