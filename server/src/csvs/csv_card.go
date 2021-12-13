package csvs

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ConfigCard struct {
	CardId       int `json:"cardId"`
	Friendliness int `json:"friendliness"` // 好感度
}

var (
	ConfigCardMap map[int]*ConfigCard
)

func GetCardConfig(cardId int) *ConfigCard {
	return ConfigCardMap[cardId]
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

			ConfigCardMap[CardId] = &ConfigCard{
				CardId, Friendliness,
			}

		}
	}

	fmt.Println(ConfigCardMap[0])
}

func init() {
	loadConfigCardCsv()
}
