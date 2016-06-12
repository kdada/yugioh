package card

import (
	"yugioh/base"
)

//怪兽卡基本信息接口
type IMonsterCard interface {
	IBaseCard
	// MonsterProperty 返回怪兽属性
	MonsterProperty() base.MonsterProperty
	// MonsterRace 返回怪兽种族
	MonsterRace() base.MonsterRace
	// MonsterSummonType 返回怪兽的召唤类型,召唤类型可能由多种类型组合而成
	// 如: MonsterSummonTypeNormal | MonsterSummonTypeEffect
	MonsterSummonType() base.MonsterSummonType
	// MonsterAttack 返回怪兽的攻击方式,攻击方式可能由多种方式组合而成
	// 如: MonsterAttackTypeNormal | MonsterAttackTypeDirect
	MonsterAttack() base.MonsterAttackType
	// MonsterDefence 返回怪兽的防御方式,目前不可能由多种防御方式组成
	MonsterDefence() base.MonsterDefenceType
	// Star 返回怪兽星级
	Star() int
	// Attack 返回怪兽攻击力
	Attack() int
	// Defence 返回怪兽防御力
	Defence() int
}
