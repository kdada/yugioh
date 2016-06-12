package card

import (
	"yugioh/base"
)

//卡片事件接口
//这些方法的执行通常与卡片环境有关
type ICardEvent interface {
	// Init 初始化卡片的自定义的字段
	Init(environment ICardEnvironment)
	// EffectSpeed 返回效果的速度
	EffectSpeed(environment ICardEnvironment, index uint) int
	// InvokeEffect 执行指定序号的效果
	InvokeEffect(environment ICardEnvironment, index uint)
	// OnInvokeEffectFiniehed 当效果执行完毕时被游戏场景调用,cancel说明效果是否成功发动
	OnInvokeEffectFiniehed(environment ICardEnvironment, index uint, cancel bool)
	// OnEnvironmentInited 注册当前卡片需要特殊的事件,在游戏环境初始化完成之后被游戏场景调用
	// 游戏开始之前,卡片的特殊效果事件都应该通过此方法被注册到游戏的事件列表中
	OnEnvironmentInited(environment ICardEnvironment)
	// OnMoving 在卡片从一个区域被移动到另一个区域之前,会被游戏场景调用
	// 可通过该方法注册或移除一些事件监听器
	// 游戏刚开始第一次调用该方法参数必然是from:SceneAreaNone to:SceneAreaDeck
	OnMoving(environment ICardEnvironment, from base.SceneArea, to base.SceneArea)
	// OnMoved 在卡片从一个区域被移动到另一个区域之后,会被场景调用
	// 可通过该方法注册或移除一些事件监听器
	// 游戏刚开始第一次调用该方法参数必然是from:SceneAreaNone to:SceneAreaDeck
	OnMoved(environment ICardEnvironment, from base.SceneArea, to base.SceneArea)
}
