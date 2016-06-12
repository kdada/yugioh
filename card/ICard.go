package card

//每张卡片都必须实现的接口
type ICard interface {
	IBaseCard  //基础卡片信息接口
	IGameCard  //游戏时卡片接口
	ICardEvent //卡片事件接口
}
