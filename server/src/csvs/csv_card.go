package csvs

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"server/utils"
	"strconv"
)

type ConfigCard struct {
	CardId       int `json:"cardId"`
	Friendliness int `json:"friendliness"` // 好感度
	Check int	`json:"check"`
}

var (
	ConfigCardMap map[int]*ConfigCard
	ConfigCardMapByRoleId map[int]*ConfigCard
)

func GetCardConfig(cardId int) *ConfigCard {
	return ConfigCardMap[cardId]
}

func GetCardConfigByRoleId(roleId int) *ConfigCard {
	return ConfigCardMapByRoleId[roleId]
}


func loadConfigCardCsv() {
	exPath, _ := os.Getwd()
	fmt.Println("path", exPath)

	fs, _ := os.Open("../../csv/Card.csv")
	result := csv.NewReader(fs)

	content, err := result.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	ConfigCardMap = make(map[int]*ConfigCard)
	for index, row := range content {
		if index > 2 {
			CardId, _ := strconv.Atoi(row[0])
			Friendliness, _ := strconv.Atoi(row[1])
			Check, _ := strconv.Atoi(row[2])
			ConfigCardMap[CardId] = &ConfigCard{
				CardId, Friendliness,Check,
			}
		}
	}

	fmt.Println(ConfigCardMap[0])
}


func init(){
	ConfigCardMap = make(map[int]*ConfigCard)
	utils.GetCsvUtilMgr().LoadCsv("Card", &ConfigCardMap)
	ConfigCardMapByRoleId = make(map[int]*ConfigCard)
	for _, v := range ConfigCardMap {
		ConfigCardMapByRoleId[v.Check] = v
	}
	return
}
