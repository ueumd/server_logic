package csvs

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ConfigPlayerLevel struct {
	PlayerLevel int `json:"PlayerLevel"`
	PlayerExp   int `json:"PlayerExp"`
	WorldLevel  int `json:"WorldLevel"`
	ChapterId   int `json:"ChapterId"`
}

var (
	ConfigPlayerLevelSlice []*ConfigPlayerLevel
)

// 加载玩家等级
func loadCsv() {
	exPath, _ := os.Getwd()
	fmt.Println("path", exPath)

	fs, _ := os.Open("./PlayerLevel.csv")
	result := csv.NewReader(fs)

	content, err := result.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	for index, row := range content {
		if index > 0 {
			PlayerLevel, _ := strconv.Atoi(row[0])
			PlayerExp, _ := strconv.Atoi(row[1])
			WorldLevel, _ := strconv.Atoi(row[2])
			ChapterId, _ := strconv.Atoi(row[3])

			ConfigPlayerLevelSlice = append(ConfigPlayerLevelSlice,
				&ConfigPlayerLevel{PlayerLevel, PlayerExp, WorldLevel, ChapterId},
			)

		}
	}
	//for index, data := range ConfigPlayerLevelSlice {
	//	fmt.Println(index, data.PlayerLevel, data.PlayerExp, data.WorldLevel, data.ChapterId)
	//}
	fmt.Println(ConfigPlayerLevelSlice[1])
}

func init() {
	loadCsv()
}

func GetNowLevelConfig(level int) *ConfigPlayerLevel {
	if level < 0 || level > len(ConfigPlayerLevelSlice) {
		return nil
	}

	return ConfigPlayerLevelSlice[level-1]
}
