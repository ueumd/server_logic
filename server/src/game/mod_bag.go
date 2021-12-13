package game

import (
	"fmt"
	"server/csvs"
)

// 背包

type ItemInfo struct {
	ItemId  int
	ItemNum int64
}
type ModBag struct {
	BagInfo map[int]*ItemInfo
}

// 增加物品 物品ID
func (self *ModBag) AddItem(itemId int, num int64, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println("物品不存在")
		return
	}
	//
	//ITEMTYPE_NORMAL = 1 // 物品
	//ITEMTYPE_ROLE   = 2 // 角色
	//ITEMTYPE_ICON   = 3 // 头像
	//ITEMTYPE_CARD   = 4 // 名片

	switch itemConfig.SortType {
	//case csvs.ITEMTYPE_NORMAL:
	//	fmt.Println("普通物品", itemConfig.ItemName)
	//	self.AddItemToBag(itemId, num)
	case csvs.ITEMTYPE_ROLE:
		fmt.Println("角色", itemConfig.ItemName)
		player.ModRole.AddItem(itemId, num)
	case csvs.ITEMTYPE_ICON:
		fmt.Println("头像", itemConfig.ItemName)
		player.ModIcon.AddItem(itemId)
	case csvs.ITEMTYPE_CARD:
		fmt.Println("名片", itemConfig.ItemName)
		player.ModCard.AddItem(itemId, 10)
	default:
		// 同普通
		self.AddItemToBag(itemId, num)
	}
}

func (self *ModBag)AddItemToBag(itemId int, num int64)  {
	_, ok := self.BagInfo[itemId]

	if ok {
		self.BagInfo[itemId].ItemNum += num

	} else {
		self.BagInfo[itemId] = &ItemInfo{
			ItemId: itemId,
			ItemNum: num,
		}
	}

	config := csvs.GetItemConfig(itemId)

	if config != nil {
		fmt.Println("获得物品", config.ItemName, "-----数量: ", num, "-----当前数量：", self.BagInfo[itemId].ItemNum)
	}
}

// 后台接口 管理者使用
func (self *ModBag)RemoveItemToBagGM(itemId int, num int64)  {
	_, ok := self.BagInfo[itemId]

	if ok {
		self.BagInfo[itemId].ItemNum -= num
	}else {
		self.BagInfo[itemId] = &ItemInfo{
			ItemId: itemId,
			ItemNum: 0-num,
		}
	}

	config := csvs.GetItemConfig(itemId)

	if config != nil {
		fmt.Println("扣除物品", config.ItemName, "-----数量: ", num, "-----当前数量：", self.BagInfo[itemId].ItemNum)
	}
}

// 正常移除物品
func (self *ModBag)RemoveItemToBag(itemId int, num int64)  {
	if !self.HasEnoughItem(itemId, num) {

		config := csvs.GetItemConfig(itemId)
		if config != nil {
			nowNum := int64(0)
			_, ok := self.BagInfo[itemId]
			if ok {
				nowNum = self.BagInfo[itemId].ItemNum
			}
			fmt.Println("物品数量不足", "-----当前数量：", nowNum)
		}
		return
	}

	_, ok := self.BagInfo[itemId]

	if ok {
		self.BagInfo[itemId].ItemNum -= num
	}else {
		self.BagInfo[itemId] = &ItemInfo{
			ItemId: itemId,
			ItemNum: 0-num,
		}
	}

	config := csvs.GetItemConfig(itemId)

	if config != nil {
		fmt.Println("扣除物品", config.ItemName, "-----数量: ", num, "-----当前数量：", self.BagInfo[itemId].ItemNum)
	}
}



func (self *ModBag)removeItem(itemId int, num int64)  {
	itemConfig := csvs.GetItemConfig(itemId)

	if itemConfig == nil {
		fmt.Println("物品不存在")
	}

	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		fmt.Println("普通物品", itemConfig.ItemName)
		self.RemoveItemToBagGM(itemId, 1)
	default:
		// 同普通
		// self.AddItemToBag(itemId, 1)
	}
}
func (self *ModBag)HasEnoughItem(itemId int, num int64) bool {
	_, ok := self.BagInfo[itemId]
	if !ok {
		return false
	} else if self.BagInfo[itemId].ItemNum < num  {
		return false
	}

	return true
}
