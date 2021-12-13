package game

import (
	"fmt"
	"server/csvs"
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
	ModBag        *ModBag    // 背包
	ModWeapon     *ModWeapon // 武器
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
	player.ModRole.RoleInfo = make(map[int]*RoleInfo)

	// 完家背包
	player.ModBag = new(ModBag)
	player.ModBag.BagInfo = make(map[int]*ItemInfo)

	// 任务
	player.ModUniqueTask = new(ModUniqueTask)
	player.ModUniqueTask.MyTaskInfo = make(map[int]*TaskInfo)
	player.ModUniqueTask.Locker = new(sync.RWMutex)

	// 武器
	player.ModWeapon = new(ModWeapon)
	player.ModWeapon.WeaponInfo = make(map[int]*Weapon)

	//****************************************
	player.ModPlayer.Icon = 0
	player.ModPlayer.PlayerLevel = 1
	player.ModPlayer.Name = "旅行者"

	// 玩家等级初始化
	player.ModPlayer.WorldLevel = 1
	player.ModPlayer.WorldLevelNow = 1
	//****************************************
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
func (self *Player) Run2() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%5 == 0 {
				// 每5s增加加1000
				self.ModBag.AddItem(1000003, 1000, self)
			} else {
				// 每1s扣300
				self.ModBag.RemoveItemToBag(1000003, 300)
			}

			//玩家报时
			// fmt.Println(time.Now().Unix())
		}
	}
}

func (self *Player) Run3() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%5 == 0 {
				// 每5s
				self.ModBag.AddItem(2000017, 7, self)
			}
		}
	}
}
func (self *Player) Run() {
	fmt.Println("===========================================")
	fmt.Println("模拟用户创建成功OK------开始测试")
	fmt.Println("===========================================")
	for {
		fmt.Println(self.ModPlayer.Name, ",欢迎来到提瓦特大陆,请选择功能：1基础信息 2背包 3(优菈UP池)模拟抽卡 4地图(未开放)")
		var modChoose int
		fmt.Scan(&modChoose)
		switch modChoose {
		case 1:
			self.HandleBase()
		}
	}
}

// 基础信息
func (self *Player) HandleBase() {
	for {
		fmt.Println("当前处于基础信息界面,请选择操作：\n0返回 \n1查询信息 \n2设置名字 \n3设置签名 \n4头像 \n5名片 \n6设置生日")

		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			self.HandleBaseGetInfo()
		}
	}
}

func (self *Player) HandleBaseGetInfo() {
	fmt.Println("名字:", self.ModPlayer.Name)
	fmt.Println("等级:", self.ModPlayer.PlayerLevel)
	fmt.Println("大世界等级:", self.ModPlayer.WorldLevelNow)
	if self.ModPlayer.Sign == "" {
		fmt.Println("签名:", "未设置")
	} else {
		fmt.Println("签名:", self.ModPlayer.Sign)
	}

	if self.ModPlayer.Icon == 0 {
		fmt.Println("头像:", "未设置")
	} else {
		fmt.Println("头像:", csvs.GetItemConfig(self.ModPlayer.Icon), self.ModPlayer.Icon)
	}

	if self.ModPlayer.Card == 0 {
		fmt.Println("名片:", "未设置")
	} else {
		fmt.Println("名片:", csvs.GetItemConfig(self.ModPlayer.Card), self.ModPlayer.Card)
	}

	if self.ModPlayer.Birth == 0 {
		fmt.Println("生日:", "未设置")
	} else {
		fmt.Println("生日:", self.ModPlayer.Birth/100, "月", self.ModPlayer.Birth%100, "日")
	}
}
