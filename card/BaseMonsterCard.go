package card

import (
	"yugioh/base"
)

//怪兽卡基本数据
//实现:IMonsterCard
type BaseMonsterCard struct {
	BaseCard                                   //基础信息
	monsterProperty    base.MonsterProperty    //怪兽属性
	monsterRace        base.MonsterRace        //怪兽种族
	monsterSummonType  base.MonsterSummonType  //怪兽召唤方式
	monsterAttackType  base.MonsterAttackType  //怪兽攻击方式
	monsterDefenceType base.MonsterDefenceType //怪兽防御方式
	star               int                     //星级
	attack             int                     //攻击力
	defence            int                     //防御力
}

// MonsterProperty 返回怪兽属性
func (this *BaseMonsterCard) MonsterProperty() base.MonsterProperty {
	return this.monsterProperty
}

// MonsterRace 返回怪兽种族
func (this *BaseMonsterCard) MonsterRace() base.MonsterRace {
	return this.monsterRace
}

// MonsterSummonType 返回怪兽的召唤类型,召唤类型可能由多种类型组合而成
// 如: MonsterSummonTypeNormal | MonsterSummonTypeEffect
func (this *BaseMonsterCard) MonsterSummonType() base.MonsterSummonType {
	return this.monsterSummonType
}

// MonsterAttack 返回怪兽的攻击方式,攻击方式可能由多种方式组合而成
// 如: MonsterAttackTypeNormal | MonsterAttackTypeDirect
func (this *BaseMonsterCard) MonsterAttack() base.MonsterAttackType {
	return this.monsterAttackType
}

// MonsterDefence 返回怪兽的防御方式,目前不可能由多种防御方式组成
func (this *BaseMonsterCard) MonsterDefence() base.MonsterDefenceType {
	return this.monsterDefenceType
}

// Star 返回怪兽星级
func (this *BaseMonsterCard) Star() int {
	return this.star
}

// Attack 返回怪兽攻击力
func (this *BaseMonsterCard) Attack() int {
	return this.attack
}

// Defence 返回怪兽防御力
func (this *BaseMonsterCard) Defence() int {
	return this.defence
}
