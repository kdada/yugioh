package card

import (
	"yugioh/base"
)

//卡片游戏时数据
//实现:IGameCard,ICardEvent
type GameCard struct {
	origin         IBaseCard      //当前卡片的原始卡片信息
	cardId         base.CardId    //卡片游戏时id,即使相同的卡,游戏时id也不同
	sceneArea      base.SceneArea //当前卡片所处的场景区域
	cardState      base.CardState //卡片状态
	ownerId        int            //拥有者id
	currentOwnerId int            //当前拥有者id
	relatedCards   []IGameCard    //关联的其他卡片
}

// Origin 返回当前卡片的原始卡片实例,即BaseMonsterCard,BaseSpellCard,BaseTrapCard
func (this *GameCard) Origin() IBaseCard {
	return this.origin
}

// CardId 返回卡片的游戏时id,id唯一,即使CardNo相同,CardId也不会相同
func (this *GameCard) CardId() base.CardId {
	return this.cardId
}

// SceneArea 返回当前卡片所处的场景位置
func (this *GameCard) SceneArea() base.SceneArea {
	return this.sceneArea
}

// SetSceneArea 设置当前卡片所处的场景位置
func (this *GameCard) SetSceneArea(area base.SceneArea) {
	this.sceneArea = area
}

// CardState 返回当前卡片的显示状态:覆盖,攻击表示等
func (this *GameCard) CardState() base.CardState {
	return this.cardState
}

// SetCardState 设置当前卡片的显示状态,只有在
func (this *GameCard) SetCardState(state base.CardState) {
	this.cardState = state
}

// OwnerId 返回当前卡片的所有者id
func (this *GameCard) OwnerId() int {
	return this.ownerId
}

// CurrentOwnerId 返回当前卡片的当前所有者id
func (this *GameCard) CurrentOwnerId() int {
	return this.currentOwnerId
}

// SetCurrentOwnerId 设置当前卡片的所有者id
func (this *GameCard) SetCurrentOwnerId(id int) {
	this.currentOwnerId = id
}

// RelatedCards 返回与当前卡片相关联的卡
func (this *GameCard) RelatedCards() []IGameCard {
	return this.relatedCards
}

// BindRelatedCards 将当前卡片和card关联,关联为单向的
// 如果要建立双向关联,还需要调用card的BindRelatedCards将当前卡片传递过去
func (this *GameCard) BindRelatedCards(card IGameCard) {
	if card != nil {
		this.relatedCards = append(this.relatedCards, card)
	}
}

// UnbindRelatedCards 取消当前卡片和card的关联,取消关联也是单向的
// 如果要取消双向关联,还需要调用card的UnbindRelatedCards将当前卡片传递过去
func (this *GameCard) UnbindRelatedCards(card IGameCard) {
	if card != nil {
		for i, c := range this.relatedCards {
			if c.CardId() == card.CardId() {
				this.relatedCards = append(this.relatedCards[:i], this.relatedCards[i+1:]...)
				return
			}
		}
	}
}

//事件部分

// Init 初始化卡片的自定义的字段
func (this *GameCard) Init(environment ICardEnvironment) {

}

// EffectSpeed 返回效果的速度 默认为1
func (this *GameCard) EffectSpeed(environment ICardEnvironment, index uint) int {
	return 1
}

// InvokeEffect 执行指定序号的效果
func (this *GameCard) InvokeEffect(environment ICardEnvironment, index uint) {
	return
}

// OnInvokeEffectFiniehed 当效果执行完毕时被游戏场景调用,cancel说明效果是否成功发动
func (this *GameCard) OnInvokeEffectFiniehed(environment ICardEnvironment, index uint, cancel bool) {
	return
}

// OnEnvironmentInited 注册当前卡片需要特殊的事件,在游戏环境初始化完成之后被游戏场景调用
// 游戏开始之前,卡片的特殊效果事件都应该通过此方法被注册到游戏的事件列表中
func (this *GameCard) OnEnvironmentInited(environment ICardEnvironment) {
	return
}

// OnMoving 在卡片从一个区域被移动到另一个区域之前,会被游戏场景调用
// 可通过该方法注册或移除一些事件监听器
// 游戏刚开始第一次调用该方法参数必然是from:SceneAreaNone to:SceneAreaDeck
func (this *GameCard) OnMoving(environment ICardEnvironment, from base.SceneArea, to base.SceneArea) {
	return
}

// OnMoved 在卡片从一个区域被移动到另一个区域之后,会被场景调用
// 可通过该方法注册或移除一些事件监听器
// 游戏刚开始第一次调用该方法参数必然是from:SceneAreaNone to:SceneAreaDeck
func (this *GameCard) OnMoved(environment ICardEnvironment, from base.SceneArea, to base.SceneArea) {
	return
}
