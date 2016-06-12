package game

//游戏通信接口
type IComunicator interface {
	// TeamCount 返回队伍数量
	TeamCount() int
	// Players 返回指定分组的玩家id
	Players(teamId int) []int
	// PostMessage 发送消息给指定玩家
	// player为玩家id
	// instruction为指令编号
	// data为数据,一般为结构体,也可能是其他类型,视具体指令而定
	PostMessage(player int, instruction uint32, data interface{})
	// GetMessage 接收玩家发送的消息,合理的情况下应该在没有消息的时候阻塞go程
	// ok返回是否接收到消息,接收到消息则为true,用于非阻塞的GetMessage方法
	GetMessage() (player int, instruction uint32, data interface{},ok bool)
}
