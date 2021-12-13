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
	// player.ModPlayer.Icon = 0
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
func (self *Player) SetBirth(birth int, player *Player) {
	self.ModPlayer.SetBirth(birth, player)
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
		case 2:
			self.HandleBagSetName()
		case 3:
			self.HandleBagSetSign()
		case 4:
			self.HandleBagSetIcon()
		case 5:
			self.HandleBagSetCard()
		case 6:
			self.HandleBagSetBirth()
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

func (self *Player) HandleBagSetName() {
	fmt.Println("请输入名字：")
	var name string
	fmt.Scan(&name)
	self.RecvSetName(name)
}

func (self *Player) HandleBagSetSign() {
	fmt.Println("请输入签名：")
	var sign string
	fmt.Scan(&sign)
	self.RecvSetSign(sign)
}

// 设置头像
func (self *Player) HandleBagSetIcon() {
	for {
		fmt.Println("当前处于基础信息--头像界面,请选择操作：0返回 1查询头像背包 2设置头像")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			self.HandleBagSetIconGetInfo()
		case 2:
			self.HandleBagSetIconSet()
		}
	}
}

func (self *Player) HandleBagSetIconGetInfo() {
	fmt.Println("当前拥有头像如下：")
	for _, v := range self.ModIcon.IconInfo {
		config := csvs.GetItemConfig(v.IconId)
		if config != nil {
			fmt.Println(config.ItemName, ":", config.ItemId)
		}
	}
}

func (self *Player) HandleBagSetIconSet() {
	fmt.Println("请输入头像id:")
	var icon int
	fmt.Scan(&icon)
	self.RecvSetIcon(icon)
}

func (self *Player) HandleBagSetCard() {
	for {
		fmt.Println("当前处于基础信息--名片界面,请选择操作：0返回1查询名片背包2设置名片")
		var action int
		fmt.Scan(&action)
		switch action {
		case 0:
			return
		case 1:
			self.HandleBagSetCardGetInfo()
		case 2:
			self.HandleBagSetCardSet()
		}
	}
}

// 名片查询
func (self *Player) HandleBagSetCardGetInfo() {
	fmt.Println("当前拥有名片如下:")
	for _, v := range self.ModCard.CardInfo {
		config := csvs.GetItemConfig(v.CardId)
		if config != nil {
			fmt.Println(config.ItemName, ":", config.ItemId)
		}
	}
}

// 名片
func (self *Player) HandleBagSetCardSet() {
	fmt.Println("请输入名片id:")
	var card int
	fmt.Scan(&card)
	self.RecvSetCard(card)
}

// 生日
func (self *Player) HandleBagSetBirth() {
	if self.ModPlayer.Birth > 0 {
		fmt.Println("已设置过生日!")
		return
	}
	fmt.Println("生日只能设置一次，请慎重填写,输入月:")
	var month, day int
	fmt.Scan(&month)
	fmt.Println("请输入日:")
	fmt.Scan(&day)
	self.ModPlayer.SetBirth(month*100+day, self)
}
