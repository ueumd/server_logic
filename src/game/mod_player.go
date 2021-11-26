package game

import (
	"fmt"
	"server/csvs"
)

// 玩家信息

type ModPlayer struct {
	UserId         int
	Icon           int
	Card           int    // 名片
	Name           string // 名字
	Sign           string // 签名
	PlayerLevel    int    // 玩家等级 由配置表导入
	PlayerExp      int    // 玩家经验 由配置表导入
	WorldLevel     int
	WorldLevelCool int64
	Birth          int
	ShowTeam       []int // 阵容模块
	ShowCard       int

	// 看不见的字段
	IsProhibit int
	IsGM       int
}

func (self *ModPlayer) SetIcon(iconId int, player *Player) {
	if !player.ModIcon.IsHasIcon(iconId) {
		// 通知客户端 非法操作
	}

	player.ModPlayer.Icon = iconId
	fmt.Println("当前图标：", player.ModPlayer.Icon)
}

func (self *ModPlayer) SetCard(cardId int, player *Player) {
	if !player.ModIcon.IsHasCard(cardId) {
		// 通知客户端 非法操作
	}
	player.ModPlayer.Icon = cardId
	fmt.Println("玩家名片：", player.ModPlayer.Icon)
}

func (self *ModPlayer) SetName(name string, player *Player) {
	// 名字较验是否合法 违禁词库
	// ...

	if GetManageBanWord().IsBanWord(name) {
		return
	}

	player.ModPlayer.Name = name
	fmt.Println("当前名字：", player.ModPlayer.Name)
}

func (self *ModPlayer) SetSign(sign string, player *Player) {
	if GetManageBanWord().IsBanWord(sign) {
		return
	}
	player.ModPlayer.Sign = sign
	fmt.Println("当家签名：", player.ModPlayer.Sign)
}

// 玩家等级
func (self *ModPlayer) AddExp(exp int, player *Player) {
	self.PlayerExp += exp

	for {
		config := csvs.GetNowLevelConfig(self.PlayerLevel)
		if config == nil {
			break
		}
		// 达到上限
		if config.PlayerExp == 0 {
			break
		}

		// 是否完成任务 才能升级
		if config.ChapterId > 0 && !player.ModUniqueTask.IsTaskFinish(config.ChapterId) {
			// 到25级不给升了，任务卡住了
			break
		}

		// 升级
		if self.PlayerExp >= config.PlayerExp {
			self.PlayerLevel += 1
			self.PlayerExp -= config.PlayerExp
		} else {
			break
		}
	}
	fmt.Println("当前等级：", self.PlayerLevel, " 当前经验：", self.PlayerExp)
}
