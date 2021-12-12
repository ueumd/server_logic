package game

import (
	"fmt"
	"sync"
	"time"
)

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
	ModPlayer     *ModPlayer     // 玩家信息
	ModIcon       *ModIcon       // 玩家Icon
	ModCard       *ModCard       // 玩家名片
	ModUniqueTask *ModUniqueTask // 任务
	ModRole       *ModeRole
	ModBag        *ModBag // 背包
}

// 测试 初始化
func NewTestPlayer() *Player {
	// 生成一个玩家
	player := new(Player)

	// 玩家信息初始化

	player.ModPlayer = new(ModPlayer)

	// 头像初始化
	player.ModIcon = new(ModIcon)
	player.ModIcon.IconInfo = make(map[int]*Icon)

	// 名片初始化
	player.ModCard = new(ModCard)
	player.ModCard.CardInfo = make(map[int]*Card)

	player.ModRole = new(ModeRole)
	player.ModBag = new(ModBag)

	// 任务
	player.ModUniqueTask = new(ModUniqueTask)
	player.ModUniqueTask.MyTaskInfo = make(map[int]*TaskInfo)
	player.ModUniqueTask.Locker = new(sync.RWMutex)

	//************************************
	// 初始化值
	player.ModPlayer.Icon = 0
	player.ModPlayer.PlayerLevel = 1

	// 玩家等级初始化
	player.ModPlayer.WorldLevel = 6
	player.ModPlayer.WorldLevelNow = 6
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

// 降低世界等级
func (self *Player) ReduceWorldLevel() {
	self.ModPlayer.ReduceWorldLevel(self)
}

// 返回等级
func (self *Player) ReturnWorldLevel() {
	self.ModPlayer.ReturnWorldLevel(self)
}

// 返回等级
func (self *Player) SetBirth(birth int) {
	self.ModPlayer.SetBirth(birth)
}

// 展示名片
func (self *Player) SetShowCard(showCard []int) {
	self.ModPlayer.SetShowCard(showCard, self)
}

// 展示阵容
func (self *Player) SetShowTeam(showTeam []int) {
	self.ModPlayer.SetShowTeam(showTeam, self)
}

func (self *Player) SetHideShowTeam(isHide int) {
	self.ModPlayer.SetHideShowTeam(isHide, self)
}

// 监听客户端给服务器发送消息
func (self *Player) Run() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			// 玩家报时
			fmt.Println(time.Now().Unix())
		}
	}
}
