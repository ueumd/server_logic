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
func (self *ModBag) AddItem(itemId int, player *Player) {
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
	case csvs.ITEMTYPE_NORMAL:
		fmt.Println("普通物品", itemConfig.ItemName)
	case csvs.ITEMTYPE_ROLE:
		fmt.Println("角色", itemConfig.ItemName)
	case csvs.ITEMTYPE_ICON:
		fmt.Println("头像", itemConfig.ItemName)
		player.ModIcon.AddItem(itemId)
	case csvs.ITEMTYPE_CARD:
		fmt.Println("名片", itemConfig.ItemName)
		player.ModCard.AddItem(itemId, 10)

	}
}
