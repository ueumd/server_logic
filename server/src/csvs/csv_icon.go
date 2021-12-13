package csvs

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ConfigIcon struct {
	IconId int `json:"IconId"`
}

var (
	ConfigIconMap map[int]*ConfigIcon
)

func GetIconConfig(iconId int) *ConfigIcon {
	return ConfigIconMap[iconId]
}

func loadIconCsv() {
	exPath, _ := os.Getwd()
	fmt.Println("path", exPath)

	fs, _ := os.Open("../../csv/Icon.csv")
	result := csv.NewReader(fs)

	content, err := result.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	ConfigIconMap = make(map[int]*ConfigIcon)
	for index, row := range content {
		if index > 2 {
			IconId, _ := strconv.Atoi(row[0])
			ConfigIconMap[IconId] = &ConfigIcon{
				IconId,
			}

		}
	}

	fmt.Println(ConfigIconMap[0])
}

func init() {
	loadIconCsv()
}
