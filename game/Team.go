package game

//队伍
type Team struct {
	TeamId  int       //队伍id
	Players []*Player //队伍成员
}

func NewTeam() *Team {
	var team = new(Team)
	team.Players = make([]*Player, 0, 1)
	return team
}

// AddPlayer 添加玩家到队伍
func (this *Team) AddPlayer(player *Player) {
	this.Players = append(this.Players, player)
}
