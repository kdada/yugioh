package game

import (
	"pkg/math/list"
	"yugioh/base"
	"yugioh/card"
)

//效果节点
type EffectNode struct {
	list.LinkedNode                 //继承自链接节点
	Card            card.ICard      //卡片
	EffectType      base.EffectType //效果类型
	EffectIndex     uint            //效果序号
	EffectSpeed     int             //效果速度
	Cancel          bool            //是否取消效果,被取消的效果无法发动
}

// NewEffectNode 创建效果节点
func NewEffectNode(card card.ICard, effectIndex uint, effectType base.EffectType, speed int) *EffectNode {
	var node = new(EffectNode)
	node.Card = card
	node.EffectIndex = effectIndex
	node.EffectSpeed = speed
	node.EffectType = effectType
	return node
}
