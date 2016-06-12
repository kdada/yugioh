package game

import (
	"pkg/dispatcher"
	"pkg/math/list"
	"strconv"
	"yugioh/base"
	"yugioh/card"
)

//游戏场景
//实现:cards.ICardEnvironment
type GameScene struct {
	GameStages                  []*GameStage                //游戏阶段列表
	TeamList                    []*Team                     //队伍列表,队伍id从0开始计数
	Players                     []*Player                   //玩家列表
	CurrentPlayer               *Player                     //当前回合所属的玩家的编号
	CurrentColumn               int                         //当前玩家在队伍中的序号
	CurrentStage                *GameStage                  //当前所处阶段
	TimeCounter                 int                         //回合计数
	FieldCardSlot               *CardSlot                   //场地卡片槽位
	EffectChain                 list.IList                  //效果连锁列表 EffectNode
	GameOver                    bool                        //标志本场游戏是否结束
	WinTeam                     *Team                       //标志获胜的队伍,如果游戏未结束,则该项为nil
	GameEventDispatcher         *dispatcher.EventDispatcher //游戏事件分发器
	Comunicator                 IComunicator                //游戏通信交换器
	ComunitationEventDispatcher *dispatcher.EventDispatcher //通信事件分发器
}

// NewGameScene 创建新的游戏场景实例
func NewGameScene() *GameScene {
	var gameScene = new(GameScene)
	//创建阶段
	gameScene.GameStages = make([]*GameStage, base.GameStageTypeEnd+1)
	var stageEventName = [...]base.EventName{
		base.EventNameStageStart,
		base.EventNameStageBeforeDraw,
		base.EventNameStageDraw,
		base.EventNameStageAfterDraw,
		base.EventNameStageBeforeMainOne,
		base.EventNameStageMainOne,
		base.EventNameStageAfterMainOne,
		base.EventNameStageBeforeBattle,
		base.EventNameStageBattle,
		base.EventNameStageAfterBattle,
		base.EventNameStageBeforeMainTwo,
		base.EventNameStageMainTwo,
		base.EventNameStageAfterMainTwo,
		base.EventNameStageEnd}
	for i := base.GameStageTypeStart; i <= base.GameStageTypeEnd; i++ {
		gameScene.GameStages[i] = NewGameStage(i, stageEventName[int(i)])
	}
	gameScene.TeamList = make([]*Team, 0, 2)
	gameScene.Players = make([]*Player, 0, 2)
	gameScene.FieldCardSlot = NewCardSlot(base.CardSlotTypeField)
	gameScene.EffectChain = list.NewList()
	gameScene.GameOver = false
	gameScene.WinTeam = nil
	gameScene.GameEventDispatcher = dispatcher.NewEventDispatcher()
	gameScene.ComunitationEventDispatcher = dispatcher.NewEventDispatcher()
	return gameScene
}

// Run 启动当前游戏场景实例,运行后会循环进行消息接收
// 每个队伍的人数必须相同
func (this *GameScene) Run(comunicator IComunicator) {
	this.Comunicator = comunicator
	for teamId := 0; teamId < this.Comunicator.TeamCount(); teamId++ {
		var team = NewTeam()
		this.TeamList = append(this.TeamList, team)
		var players = this.Comunicator.Players(teamId)
		for _, playerId := range players {
			var player = NewPlayer(team, playerId)
			team.AddPlayer(player)
			this.Players = append(this.Players, player)
		}
	}
	for !this.GameOver {
		var playerId, instruction, msg, ok = this.Comunicator.GetMessage()
		if ok {
			this.ComunitationEventDispatcher.DispatchEvent(strconv.Itoa(int(instruction)), playerId, msg)
		}
	}
}

// RegisterEvents 注册事件,包括基本事件,卡片事件
func (this *GameScene) RegisterEvents() {
	//注册阶段事件
	this.RegisterEvent(NewStageAutoChangedEvent(base.EventNameStageStart))
	this.RegisterEvent(NewStageAutoChangedEvent(base.EventNameStageBeforeDraw))
	this.RegisterEvent(NewStageAutoChangedEvent(base.EventNameStageDraw))
	this.RegisterEvent(NewStageAutoChangedEvent(base.EventNameStageAfterDraw))
	this.RegisterEvent(NewStageAutoChangedEvent(base.EventNameStageBeforeMainOne))
	this.RegisterEvent(NewStageChangedEvent(base.EventNameStageMainOne))
	this.RegisterEvent(NewStageAutoChangedEvent(base.EventNameStageAfterMainOne))
	this.RegisterEvent(NewStageAutoChangedEvent(base.EventNameStageBeforeBattle))
	this.RegisterEvent(NewStageChangedEvent(base.EventNameStageBattle))
	this.RegisterEvent(NewStageAutoChangedEvent(base.EventNameStageAfterBattle))
	this.RegisterEvent(NewStageAutoChangedEvent(base.EventNameStageBeforeMainTwo))
	this.RegisterEvent(NewStageAutoChangedEvent(base.EventNameStageMainTwo))
	this.RegisterEvent(NewStageChangedEvent(base.EventNameStageAfterMainTwo))
	this.RegisterEvent(NewStageAutoChangedEvent(base.EventNameStageEnd))
}

// RegisterListeners 注册监听器,仅包括基本监听器
func (this *GameScene) RegisterListeners() {
	//注册事件监听器
	this.RegisterListener(base.EventNameStageDraw, func(environment card.ICardEnvironment) {
		environment.DrawCard(environment.CurrentPlayerId(), 1)
	})
}

// GetPlayerById 通过id查找到玩家
func (this *GameScene) GetPlayerById(id int) (*Player, bool) {
	for _, player := range this.Players {
		if player.PlayerId == id {
			return player, true
		}
	}
	return nil, false
}

//
//功能方法,这些方法将在游戏运行过程中被使用
//

// Time 返回当前回合数
func (this *GameScene) Time() int {
	return this.TimeCounter
}

// Stage 返回当前回合所处的阶段
func (this *GameScene) Stage() base.GameStageType {
	return this.CurrentStage.GameStageType
}

// CurrentPlayerId 返回当前回合的玩家id
func (this *GameScene) CurrentPlayerId() int {
	return this.CurrentPlayer.PlayerId
}

// NextStage 切换到下一阶段没有被跳过的阶段
func (this *GameScene) NextStage() {
	if (this.CurrentStage != nil) && (this.CurrentStage.GameStageType < base.GameStageTypeEnd) {
		for _, stage := range this.GameStages {
			if (stage.GameStageType > this.CurrentStage.GameStageType) && (!stage.Skip) {
				this.GoToStage(stage.GameStageType)
			}
		}
	} else {
		//如果无法切换到下一个阶段,则直接跳过当前回合
		this.NextRange()
	}

}

// GoToStage 切换到指定的阶段
func (this *GameScene) GoToStage(stageType base.GameStageType) {
	for _, stage := range this.GameStages {
		if stage.GameStageType == stageType {
			this.CurrentStage = stage
			this.DisptachEvent(stage.GameStageName)
		}
	}
}

// NextRange 进入下一回合,如果当前处于第0回合,则自动初始化回合信息
func (this *GameScene) NextRange() {
	this.TimeCounter++
	if this.TimeCounter <= 1 {
		//游戏刚开始
		this.CurrentColumn = 0
		this.CurrentPlayer = this.TeamList[0].Players[0]
	} else {
		var currentTeamId = this.CurrentPlayer.Team.TeamId
		var teamCount = len(this.TeamList)
		if currentTeamId >= (teamCount - 1) {
			//如果超出则选择下一个队伍列
			this.CurrentColumn++
			this.CurrentColumn = this.CurrentColumn % len(this.TeamList[0].Players)
		}
		currentTeamId++
		currentTeamId = currentTeamId % teamCount
		this.CurrentPlayer = this.TeamList[currentTeamId].Players[this.CurrentColumn]
	}
	//重置回合阶段
	for _, stage := range this.GameStages {
		stage.Resume()
	}
	//进入第一阶段
	this.GoToStage(base.GameStageTypeStart)
}

// DrawCard 指定player抽count张卡
func (this *GameScene) DrawCard(playerId, count int) {
	var player, ok = this.GetPlayerById(playerId)
	if ok {
		for i := 0; i < count; i++ {
			var card, ok = player.PopDeckCard()
			if ok {
				//添加到手牌
				player.AddHandCard(card)
			} else {
				// 如果玩家卡组里没有足够的卡牌了,则该玩家失败
				player.GameOver = true
				//this.TestWin()
			}
		}
	}
}

// AddHealth 增加玩家生命值,count为负数则为减
func (this *GameScene) AddHealth(playerId, count int) {
	var player, ok = this.GetPlayerById(playerId)
	if ok {
		player.Health += count
		if player.Health <= 0 {
			player.Health = 0
			player.GameOver = true
			//this.TestWin()
		}
	}
}

// AddValidEffect 给当前卡片所属玩家添加可发动效果
func (this *GameScene) AddValidEffect(card card.ICard, effectIndex uint) {
	var player, ok = this.GetPlayerById(card.CurrentOwnerId())
	if ok {
		var node = NewEffectNode(card, effectIndex, card.EffectType(effectIndex), card.EffectSpeed(this, effectIndex))
		player.AddValidEffect(node)
	}
}

// InvokeEffectIndex 将效果加入连锁列表,等待发动卡片指定效果
func (this *GameScene) AddEffectToChain(card card.ICard, effectIndex uint) {
	var node = NewEffectNode(card, effectIndex, card.EffectType(effectIndex), card.EffectSpeed(this, effectIndex))
	this.EffectChain.Push(node)
}

// InvokeEffect 发动效果
func (this *GameScene) InvokeEffect(card card.ICard, effectIndex uint) {
	card.InvokeEffect(this, effectIndex)
}

// CanSummon 检查场景是否允许进行指定类型的召唤
func (this *GameScene) CanSummon(summonType base.MonsterSummonType) bool {
	return true
}

// MoveCard 移动卡片位置,只有to为怪兽或魔法陷阱区域state才有效
func (this *GameScene) MoveCard(card card.ICard, to base.SceneArea, state base.CardState) {
	var preArea = card.SceneArea()
	card.OnMoving(this, preArea, to)
	this.DisptachEvent(base.EventNameCardMoving, preArea, to)
	this.DisptachEvent(base.EventNameCardMoved, preArea, to)
	card.OnMoved(this, preArea, to)
	card.SetSceneArea(to)
	if to == base.SceneAreaMonster || to == base.SceneAreaEffect {
		card.SetCardState(state)
	} else {
		card.SetCardState(base.CardStateNone)
	}
}

// OfferingMonsterCards 返回指定玩家可以作为祭品的怪兽卡片
func (this *GameScene) OfferingMonsterCards(playerId int) (cards []card.ICard,ok bool) {
	var player, ok = this.GetPlayerById(playerId)
	cards = make([]card.ICard,0,5)
	if ok {
		var team = player.Team
		for _, p := range team.Players {
			for _,slot := range p.MonsterArea {
				cards = append(cards,slot.Card)
			}
		}
	}

}

// SelectCards 让玩家选择从卡片中选择指定数量的卡片
func (this *GameScene) SelectCards(playerId int, cards []card.ICard, count int) []card.ICard {
	//////
	return cards
}

// TestWin 检查是否存在胜利队伍
func (this *GameScene) TestWin() {
	var aliveTeamId = -1
	var aliveCount = 0
	for teamId, team := range this.TeamList {
		var alivePlayerCount = 0
		for _, player := range team.Players {
			if !player.GameOver {
				alivePlayerCount++
			}
		}
		if alivePlayerCount > 0 {
			aliveTeamId = teamId
			aliveCount++
		}
	}
	if aliveCount == 0 {
		//平局==
		this.GameOver = true
		this.WinTeam = nil
	}
	if aliveCount == 1 {
		//有胜利者==
		this.GameOver = true
		this.WinTeam = this.TeamList[aliveTeamId]
	}
}

// DisptachEvent 分发事件
func (this *GameScene) DisptachEvent(eventName base.EventName, params ...interface{}) {
	this.GameEventDispatcher.DispatchEvent(string(eventName), params...)
}

// RegisterEvent 注册事件
func (this *GameScene) RegisterEvent(event dispatcher.IEvent) {
	this.GameEventDispatcher.RegisterEvent(event)
}

// RegisterListener 注册事件监听器
func (this *GameScene) RegisterListener(eventName base.EventName, listener interface{}) dispatcher.ListenerId {
	return this.GameEventDispatcher.RegisterEventListener(string(eventName), listener)
}

// RemoveListener 移除事件监听器
func (this *GameScene) RemoveListener(eventName base.EventName, listenerId dispatcher.ListenerId) {
	this.GameEventDispatcher.RemoveEventListenerById(string(eventName), listenerId)
}
