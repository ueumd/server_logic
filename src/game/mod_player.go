package game

import (
	"fmt"
	"server/csvs"
	"time"
)

// 玩家信息
type ModPlayer struct {
	UserId         int    // 唯id
	Icon           int    // 头像
	Card           int    // 名片
	Name           string // 名字
	Sign           string // 签名
	PlayerLevel    int    // 玩家等级 	由配置表导入
	PlayerExp      int    // 阅历（经验） 由配置表导入
	WorldLevel     int    // 大世界等级
	WorldLevelNow  int    // 大世界等级(当前)
	WorldLevelCool int64  // 操作大世界等级冷确时间
	Birth          int    // 生日
	ShowTeam       []int  // 展示阵容
	ShowCard       int    // 展示名片

	// 游戏中看不见的字段
	IsProhibit int // 封禁状态
	IsGM       int // GM帐号标志
}

func (self *ModPlayer) SetIcon(iconId int, player *Player) {
	if !player.ModIcon.IsHasIcon(iconId) {
		// 通知客户端 非法操作
		return
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

// 降低世界等级
func (self *ModPlayer) ReduceWorldLevel(player *Player) {
	// 达到5级才可以降低等级
	if self.WorldLevel < csvs.REDUCE_WORLD_LEVEL_START {
		fmt.Println("操作失败： ---当前世界等级：", self.WorldLevel)
		return
	}

	if self.WorldLevel-self.WorldLevelNow >= csvs.REDUCE_WORLD_LEVEL_MAX {
		fmt.Println("操作失败： ---当前世界等级：", self.WorldLevel, "---真实世界等级：", self.WorldLevelNow)
		return
	}

	if time.Now().Unix() < self.WorldLevelCool {
		fmt.Println("操作失败： ---冷确中")
		return
	}

	self.WorldLevelNow -= 1
	self.WorldLevelCool = time.Now().Unix() + csvs.REDUCE_WORLD_LEVEL_COOL_TIME

	fmt.Println("操作成功： ---当前世界等级：", self.WorldLevel, "---真实世界等级：", self.WorldLevelNow)
	return
}

// 返回等级
func (self *ModPlayer) ReturnWorldLevel(player *Player) {
	if self.WorldLevelNow == self.WorldLevel {
		fmt.Println("操作失败： ---当前世界等级：", self.WorldLevel, "---真实世界等级：", self.WorldLevelNow)
		return
	}

	if time.Now().Unix() < self.WorldLevelCool {
		fmt.Println("操作失败： ---冷确中")
		return
	}

	self.WorldLevelNow += 1
	self.WorldLevelCool = time.Now().Unix() + csvs.REDUCE_WORLD_LEVEL_COOL_TIME

	fmt.Println("操作成功： ---当前世界等级：", self.WorldLevel, "---真实世界等级：", self.WorldLevelNow)
	return
}

// 设置生日
func (self *ModPlayer) SetBirth(birth int) {

	month := birth / 100
	day := birth % 100

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day <= 0 || day > 31 {
			fmt.Println(month, "月没有", day, "日！")
			return
		}
	case 4,6, 9, 11:
		if day <= 0 || day > 30 {
			fmt.Println(month, "月没有", day, "日！")
			return
		}
	case 2:
		if day <= 0 || day > 29 {
			fmt.Println(month, "月没有", day, "日！")
			return
		}
	default:
		fmt.Println( "没有", month, "月!")
		return
	}

	self.Birth = birth
	fmt.Println("生日设置成功：", month, "月", day, "日")

	if self.IsBirthDay() {
		fmt.Println("今天是你的生日， 生日快乐")
	}else {
		fmt.Println("期待你的生日到来~~~")
	}
}

//  东八区时间
func (self *ModPlayer) IsBirthDay() bool{
   month := time.Now().Month()
   day := time.Now().Day()

   if int(month) == self.Birth/100 && day == self.Birth%100 {
	   return  true
   }
   return false
}