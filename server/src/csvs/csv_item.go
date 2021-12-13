package csvs

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 对应表里SortType
const (
	ITEMTYPE_NORMAL = 1 // 物品
	ITEMTYPE_ROLE   = 2 // 角色
	ITEMTYPE_ICON   = 3 // 头像
	ITEMTYPE_CARD   = 4 // 名片
)

type ConfigItem struct {
	ItemId   int    `json:"itemId"`
	SortType int    `json:"sortType"`
	ItemName string `json:"itemName"`
}

var (
	ConfigItemMap map[int]*ConfigItem
)

func GetItemConfig(itemId int) *ConfigItem {
	return ConfigItemMap[itemId]
}
func loadConfigItemCsv() {
	exPath, _ := os.Getwd()
	fmt.Println("path", exPath)

	fs, _ := os.Open("../../csv/Item.csv")
	result := csv.NewReader(fs)

	content, err := result.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	ConfigItemMap = make(map[int]*ConfigItem)
	for index, row := range content {
		if index > 2 {
			ItemId, _ := strconv.Atoi(row[0])
			SortType, _ := strconv.Atoi(row[1])
			ItemName := row[2]

			ConfigItemMap[ItemId] = &ConfigItem{
				ItemId, SortType, ItemName,
			}

		}
	}

	fmt.Println(ConfigItemMap[0])
}

func init() {
	loadConfigItemCsv()
}
