package base

//卡片编号
type CardNo uint

//卡片游戏时id
type CardId uint

//卡片类型
type CardType uint8

const (
	CardTypeNone    CardType = 0               //无
	CardTypeMonster CardType = 1 << (iota - 1) //怪兽卡
	CardTypeSpell                              //魔法卡
	CardTypeTrap                               //陷阱卡
)

//怪兽属性
type MonsterProperty uint8

const (
	MonsterPropertyNone   MonsterProperty = 0               //无
	MonsterPropertyWind   MonsterProperty = 1 << (iota - 1) //风
	MonsterPropertyFire                                     //炎
	MonsterPropertyWater                                    //水
	MonsterPropertyEarth                                    //地
	MonsterPropertyLight                                    //光
	MonsterPropertyDark                                     //暗
	MonsterPropertyDivine                                   //神
)

//怪兽种族
type MonsterRace uint32

const (
	MonsterRaceNone         MonsterRace = 0               //无
	MonsterRaceDragon       MonsterRace = 1 << (iota - 1) //龙族
	MonsterRaceSpellcaster                                //魔法使族
	MonsterRaceZombie                                     //不死族
	MonsterRaceWarrior                                    //战士族
	MonsterRaceBeast                                      //兽族
	MonsterRaceBeastWarrior                               //兽战士族
	MonsterRaceWingedBeast                                //鸟兽族
	MonsterRaceFiend                                      //恶魔族
	MonsterRaceFairy                                      //天使族
	MonsterRaceInsect                                     //昆虫族
	MonsterRaceDinosaur                                   //恐龙族
	MonsterRaceReptile                                    //爬虫族
	MonsterRaceFish                                       //鱼族
	MonsterRaceSeaSerpent                                 //海龙族
	MonsterRaceMachine                                    //机械族
	MonsterRaceThunder                                    //雷族
	MonsterRaceAqua                                       //水族
	MonsterRacePyro                                       //炎族
	MonsterRaceRock                                       //岩石族
	MonsterRacePlant                                      //植物族
	MonsterRacePsychic                                    //念动力族
	MonsterRaceDivineBeast                                //幻神兽
	MonsterRaceCreatorGod                                 //创造神
	MonsterRaceWyrm                                       //幻龙族
)

//怪兽效果
//type MonsterSkill uint8

//const (
//	MonsterSkillNone   MonsterSkill = iota //无效果
//	MonsterSkillEffect                     //有效果
//)

//怪兽召唤方式
type MonsterSummonType uint8

const (
	MonsterSummonTypeNormal   MonsterSummonType = 1 << iota //通常召唤
	MonsterSummonTypeEffect                                 //特殊召唤
	MonsterSummonTypeFusion                                 //融合召唤
	MonsterSummonTypeRitual                                 //仪式召唤
	MonsterSummonTypeSynchro                                //同调召唤
	MonsterSummonTypeXyz                                    //XYZ召唤(超量召唤)
	MonsterSummonTypePendulum                               //灵摆召唤
	MonsterSummonTypeToken                                  //代币召唤(例如因替罪羊的效果进行的召唤)
)

//怪兽攻击类型
type MonsterAttackType uint8

const (
	MonsterAttackTypeNone   MonsterAttackType = 1 << iota //无法攻击
	MonsterAttackTypeNormal                               //普通攻击方式
	MonsterAttackTypeDirect                               //直接攻击方式
)

//怪兽防御方式
type MonsterDefenceType uint8

const (
	MonsterDefenceTypeNone   MonsterDefenceType = 1 << iota //无法被攻击
	MonsterDefenceTypeNormal                                //一般防御方式
	MonsterDefenceTypeTaunt                                 //嘲讽,必须被优先攻击
)

//魔法卡类型
type SpellCardType uint8

const (
	SpellCardTypeNormal     SpellCardType = (1 << 6) | (1 << iota) //普通魔法
	SpellCardTypeContinuous                                        //永续魔法
	SpellCardTypeEquip                                             //装备魔法
	SpellCardTypeField                                             //场地魔法
	SpellCardTypeRitual                                            //仪式魔法
	SpellCardTypeQuick      SpellCardType = (2 << 6) | (1 << iota) //速攻魔法
)

//陷阱卡类型
type TrapCardType uint8

const (
	TrapCardTypeNormal     TrapCardType = (2 << 6) | (1 << iota) //普通陷阱
	TrapCardTypeContinuous                                       //永续陷阱
	TrapCardTypeCounter    TrapCardType = (3 << 6) | (1 << iota) //反击陷阱
)

//场地区域
type SceneArea uint16

const (
	SceneAreaNone    SceneArea = 0               //无
	SceneAreaDeck    SceneArea = 1 << (iota - 1) //牌组
	SceneAreaHand                                //手牌
	SceneAreaMonster                             //怪兽区
	SceneAreaEffect                              //魔法陷阱区
	SceneAreaField                               //场地区
	SceneAreaGrave                               //墓地
	SceneAreaExtra                               //额外卡组区
	SceneAreaRemove                              //移除区
)

//卡片状态
type CardState uint8

const (
	CardStateNone            CardState = 0                                    //无
	CardStateFront           CardState = 1 << (iota - 1)                      //卡片正面
	CardStateBack                                                             //卡片反面
	CardStateHorizontal                                                       //卡片水平
	CardStateVertical                                                         //卡片竖直
	CardStateFrontHorizontal CardState = CardStateFront | CardStateHorizontal //怪兽处于正面守备状态
	CardStateFrontVertical   CardState = CardStateFront | CardStateVertical   //怪兽攻击状态或魔法陷阱发动状态
	CardStateBackHorizontal  CardState = CardStateBack | CardStateHorizontal  //怪兽处于背面守备状态
	CardStateBackVertical    CardState = CardStateBack | CardStateVertical    //魔法陷阱覆盖状态
)
