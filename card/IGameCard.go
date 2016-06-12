package card

import (
	"yugioh/base"
)

//游戏时效果接口
type IGameCard interface {
	// Origin 返回当前卡片的原始卡片实例,即BaseMonsterCard,BaseSpellCard,BaseTrapCard
	Origin() IBaseCard
	// CardId 返回卡片的游戏时id,id唯一,即使CardNo相同,CardId也不会相同
	CardId() base.CardId
	// SceneArea 返回当前卡片所处的场景位置
	SceneArea() base.SceneArea
	// SetSceneArea 设置当前卡片所处的场景位置
	SetSceneArea(area base.SceneArea)
	// CardState 返回当前卡片的显示状态:覆盖,攻击表示等
	CardState() base.CardState
	// SetCardState 设置当前卡片的显示状态,只有在
	SetCardState(state base.CardState)
	// OwnerId 返回当前卡片的所有者id
	OwnerId() int
	// CurrentOwnerId 返回当前卡片的当前所有者id
	CurrentOwnerId() int
	// SetCurrentOwnerId 设置当前卡片的所有者id
	SetCurrentOwnerId(id int)
	// RelatedCards 返回与当前卡片相关联的卡
	RelatedCards() []IGameCard
	// BindRelatedCards 将当前卡片和card关联,关联为单向的
	// 如果要建立双向关联,还需要调用card的BindRelatedCards将当前卡片传递过去
	BindRelatedCards(card IGameCard)
	// UnbindRelatedCards 取消当前卡片和card的关联,取消关联也是单向的
	// 如果要取消双向关联,还需要调用card的UnbindRelatedCards将当前卡片传递过去
	UnbindRelatedCards(card IGameCard)
}
