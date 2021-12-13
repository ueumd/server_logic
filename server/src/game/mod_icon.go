package game

import (
	"fmt"
	"server/csvs"
)

type Icon struct {
	IconId int
}

// 玩家头像模块
type ModIcon struct {
	IconInfo map[int]*Icon
}

func (self *ModIcon) IsHasIcon(iconId int) bool {
	_, ok := self.IconInfo[iconId]
	return ok
}

// 头像
func (self *ModIcon) AddItem(itemId int) {
	_, ok := self.IconInfo[itemId]

	if ok {
		fmt.Println("已存在头像", itemId)
		return
	}

	config := csvs.GetIconConfig(itemId)

	if config == nil {
		fmt.Println("非法头像", itemId)
		return
	}

	self.IconInfo[itemId] = &Icon{
		IconId: itemId,
	}

	fmt.Println("获得头像", itemId)
}
