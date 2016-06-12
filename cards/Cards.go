package cards

import (
	"reflect"
	"yugioh/base"
	"yugioh/card"
)

//卡表
var cardTable = make(map[base.CardNo]*cardPackage, 100)

//卡片组合
type cardPackage struct {
	baseCard card.IBaseCard //卡片基本信息
	cardType reflect.Type   //卡片类型
}

//卡片基础信息
type cardBaseInfo struct {
	cardNo             base.CardNo             //卡片编号,同样的卡编号相同
	cardName           string                  //卡片名称
	cardType           base.CardType           //卡片类型
	effectCount        uint8                   //当前卡片拥有的效果数量
	spellCardType      base.SpellCardType      //魔法卡类型,仅在卡片类型为base.CardTypeSpell时才需要设置
	trapCardType       base.TrapCardType       //陷阱卡类型,仅在卡片类型为base.CardTypeTrap时才需要设置
	monsterProperty    base.MonsterProperty    //怪兽属性,仅在卡片类型为base.CardTypeMonster时才需要设置
	monsterRace        base.MonsterRace        //怪兽种族,仅在卡片类型为base.CardTypeMonster时才需要设置
	monsterSummonType  base.MonsterSummonType  //怪兽召唤方式,仅在卡片类型为base.CardTypeMonster时才需要设置
	monsterAttackType  base.MonsterAttackType  //怪兽攻击方式,仅在卡片类型为base.CardTypeMonster时才需要设置
	monsterDefenceType base.MonsterDefenceType //怪兽防御方式,仅在卡片类型为base.CardTypeMonster时才需要设置
	star               int                     //怪兽星级,仅在卡片类型为base.CardTypeMonster时才需要设置
	attack             int                     //攻击力,仅在卡片类型为base.CardTypeMonster时才需要设置
	defence            int                     //防御力,仅在卡片类型为base.CardTypeMonster时才需要设置
}

// registerCard 使用info注册卡片信息到卡表
// newCard必须继承自card.MonsterCard或card.SpellCard或card.TrapCard
func registerCard(info *cardBaseInfo, newCard card.ICard) {
	if (info != nil) && (newCard != nil) {
		var baseCard card.IBaseCard
		switch info.cardType {
		case base.CardTypeMonster:
			//创建基础怪兽卡信息
			baseCard = card.NewBaseMonsterCard(info.cardNo, info.cardName,
				info.cardType, info.effectCount, info.monsterProperty,
				info.monsterRace, info.monsterSummonType,
				info.monsterAttackType, info.monsterDefenceType,info.star,
				info.attack, info.defence)
		case base.CardTypeSpell:
			//创建基础魔法卡信息
			baseCard = card.NewBaseSpellCard(info.cardNo, info.cardName,
				info.cardType, info.effectCount, info.spellCardType)
		case base.CardTypeTrap:
			//创建基础魔法卡信息
			baseCard = card.NewBaseTrapCard(info.cardNo, info.cardName,
				info.cardType, info.effectCount, info.trapCardType)
		default:
			return
		}
		var cp = new(cardPackage)
		cp.baseCard = baseCard
		cp.cardType = reflect.TypeOf(newCard).Elem()
		cardTable[baseCard.CardNo()] = cp
	}
}

// CreateCardInstance 创建指定编号的卡片的游戏时实例
// 实例仅仅包含基础信息,创建之后还需要额外调用实例的Init方法
// cardNo:卡片编号
// cardId:分配给卡片实例的id
// ownerId:拥有者
func CreateCardInstance(cardNo base.CardNo, cardId base.CardId, ownerId int) (card.ICard, bool) {
	var cp, ok = cardTable[cardNo]
	if ok {
		var c, ok = reflect.New(cp.cardType).Interface().(card.ICard)
		if ok {
			ok = card.SetCardBaseInfo(c, cp.baseCard, cardId, ownerId)
			if ok {
				return c, true
			}
		}

	}
	return nil, false
}
