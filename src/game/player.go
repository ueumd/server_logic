package game

type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
	ModCard   *ModCard
}

// 测试
func NewTestPlayer() *Player {
	// 生成一个玩家
	player := new(Player)

	// 玩家信息初始化
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModCard = new(ModCard)
	player.ModPlayer.Icon = 0

	player.ModPlayer.PlayerLevel = 1

	return player
}

// 对外接口
func (self *Player) RecvSetIcon(iconId int) {
	self.ModPlayer.SetIcon(iconId, self)
}

func (self *Player) RecvSetCard(cardId int) {
	self.ModPlayer.SetCard(cardId, self)
}

func (self *Player) RecvSetName(name string) {
	self.ModPlayer.SetName(name, self)
}

func (self *Player) RecvSetSign(name string) {
	self.ModPlayer.SetSign(name, self)
}
