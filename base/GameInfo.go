package base

//游戏阶段类型
type GameStageType uint8

const (
	GameStageTypeStart         GameStageType = iota //开始阶段
	GameStageTypeBeforeDraw                         //进入抽牌阶段之前
	GameStageTypeDraw                               //抽牌阶段
	GameStageTypeAfterDraw                          //抽牌阶段结束
	GameStageTypeBeforeMainOne                      //进入主要阶段1之前
	GameStageTypeMainOne                            //主要阶段1
	GameStageTypeAfterMainOne                       //主要阶段1结束
	GameStageTypeBeforeBattle                       //进入战斗阶段之前
	GameStageTypeBattle                             //战斗阶段
	GameStageTypeAfterBattle                        //战斗阶段结束
	GameStageTypeBeforeMainTwo                      //进入主要阶段2之前
	GameStageTypeMainTwo                            //主要阶段2
	GameStageTypeAfterMainTwo                       //主要阶段2结束
	GameStageTypeEnd                                //结束阶段
)

//主动发动的效果类型
type EffectType uint8

const (
	EffectTypeNone        EffectType = iota //无
	EffectTypeAuto                          //自动效果,自动发动的效果应该被标记为该类型
	EffectTypeSummon                        //攻击召唤怪兽效果
	EffectTypeSummonCover                   //里侧防御召唤怪兽效果
	EffectTypeSummonFront                   //表侧防御召唤怪兽效果
	EffectTypeLaunch                        //发动卡片效果
	EffectTypeCover                         //覆盖卡片效果
)

//卡槽位类型
type CardSlotType uint8

const (
	CardSlotTypeMonster CardSlotType = 1 << iota //怪兽卡槽
	CardSlotTypeEffect                           //魔法陷阱卡槽
	CardSlotTypeField                            //场地卡槽
)

//事件名称
type EventName string

const (
	EventNameStageStart         EventName = "EventNameStageStart"         //开始阶段事件名称
	EventNameStageBeforeDraw    EventName = "EventNameStageBeforeDraw"    //抽牌阶段之前
	EventNameStageDraw          EventName = "EventNameStageDraw"          //抽牌阶段
	EventNameStageAfterDraw     EventName = "EventNameStageAfterDraw"     //抽牌阶段结束
	EventNameStageBeforeMainOne EventName = "EventNameStageBeforeMainOne" //进入主要阶段1之前
	EventNameStageMainOne       EventName = "EventNameStageMainOne"       //主要阶段1
	EventNameStageAfterMainOne  EventName = "EventNameStageAfterMainOne"  //主要阶段1结束
	EventNameStageBeforeBattle  EventName = "EventNameStageBeforeBattle"  //进入战斗阶段之前
	EventNameStageBattle        EventName = "EventNameStageBattle"        //战斗阶段
	EventNameStageAfterBattle   EventName = "EventNameStageAfterBattle"   //战斗阶段结束
	EventNameStageBeforeMainTwo EventName = "EventNameStageBeforeMainTwo" //进入主要阶段2之前
	EventNameStageMainTwo       EventName = "EventNameStageMainTwo"       //主要阶段2
	EventNameStageAfterMainTwo  EventName = "EventNameStageAfterMainTwo"  //主要阶段2结束
	EventNameStageEnd           EventName = "EventNameStageEnd"           //结束阶段
	EventNameCardMoving         EventName = "EventNameCardMoving"         //卡片移动之前
	EventNameCardMoved          EventName = "EventNameCardMoved"          //卡片移动之后
)
