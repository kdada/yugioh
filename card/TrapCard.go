package card

//游戏时陷阱卡
type TrapCard struct {
	BaseTrapCard //基础陷阱卡信息
	GameCard     //游戏时卡片信息

}

// EffectSpeed 返回效果的速度
func (this *TrapCard) EffectSpeed(environment ICardEnvironment, index uint) int {
	return (int(this.trapCardType) >> 6) & 0x3
}
