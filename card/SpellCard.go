package card

//游戏时魔法卡
type SpellCard struct {
	BaseSpellCard //基础魔法卡信息
	GameCard      //游戏时卡片信息
}

// EffectSpeed 返回效果的速度
func (this *SpellCard) EffectSpeed(environment ICardEnvironment, index uint) int {
	return (int(this.spellCardType) >> 6) & 0x3
}
