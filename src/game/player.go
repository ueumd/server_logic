package game

import "sync"

/**
实例化一个玩家信息
 */

// 任务状态
const (
	TASK_STATE_INIT   = 0
	TASK_STATE_DOING  = 1 // 进入副本
	TASK_STATE_FINISH = 2 // 完成
)

// 玩家信息装组
type Player struct {
	ModPlayer     *ModPlayer	// 玩家信息
	ModIcon       *ModIcon		// 玩家Icon
	ModCard       *ModCard		// 玩家名片
	ModUniqueTask *ModUniqueTask // 任务
}

// 测试 初始化
func NewTestPlayer() *Player {
	// 生成一个玩家
	player := new(Player)

	// 玩家信息初始化
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModCard = new(ModCard)
	player.ModUniqueTask = new(ModUniqueTask)
	player.ModUniqueTask.MyTaskInfo = make(map[int]*TaskInfo)
	player.ModUniqueTask.Locker = new(sync.RWMutex)

	//************************************
	// 初始化值
	player.ModPlayer.Icon = 0
	player.ModPlayer.PlayerLevel = 1
	//************************************

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
