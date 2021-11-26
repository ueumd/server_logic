package csvs

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ConfigUniqueTask struct {
	TaskId	int `json:"TaskId"`
	SortType	int `json:"SortType"`
	OpenLevel	int `json:"OpenLevel"`
	TaskType	int`json:"TaskType"`
	Condition int	`json:"Condition"`
}


var ConfigUniqueTaskMap map[int]*ConfigUniqueTask

// 任务
func loadUniqueTaskCsv() {
	exPath, _ := os.Getwd()
	fmt.Println("path", exPath)

	fs, _ := os.Open("../csv/UniqueTask.csv")
	result := csv.NewReader(fs)

	content, err := result.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}

	ConfigUniqueTaskMap = make(map[int]*ConfigUniqueTask)

	for index, row := range content {
		if index > 1 {
			TaskId, _ := strconv.Atoi(row[0])
			SortType, _ := strconv.Atoi(row[1])
			OpenLevel, _ := strconv.Atoi(row[2])
			TaskType, _ := strconv.Atoi(row[3])
			Condition, _ := strconv.Atoi(row[4])

			ConfigUniqueTaskMap[TaskId] = &ConfigUniqueTask {
				TaskId,
				SortType,
				OpenLevel,
				TaskType,
				Condition,
			}
		}
	}
	// fmt.Println(ConfigUniqueTaskMap)
}

func init()  {
	loadUniqueTaskCsv()
	return
}