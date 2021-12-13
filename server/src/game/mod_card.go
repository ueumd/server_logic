package game

import (
	"fmt"
	"server/csvs"
)

type Card struct {
	CardId int
}

// 玩家名片
type ModCard struct {
	CardInfo map[int]*Card
}

func (self *ModCard) IsHasCard(cardId int) bool {
	_, ok := self.CardInfo[cardId]
	return ok
}

// 添加名片
// friendliness 好感度 10
func (self *ModCard) AddItem(itemId int, friendliness int) {
	_, ok := self.CardInfo[itemId]

	if ok {
		fmt.Println("名片已存在", itemId)
		return
	}

	config := csvs.GetCardConfig(itemId)

	if config == nil {
		fmt.Println("非法名乍", itemId)
		return
	}

	// 如果小于配置
	if friendliness < config.Friendliness {
		fmt.Println("好感度不足", itemId)
		return
	}

	self.CardInfo[itemId] = &Card{
		CardId: itemId,
	}

	fmt.Println("获得名片", itemId)
}
