package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type CsvTable struct {
	FileName string
	Records  []CsvRecord
}

type CsvRecord struct {
	Record map[string]string
}

func (c *CsvRecord) GetInt(field string) int {
	var r int
	var err error
	if r, err = strconv.Atoi(c.Record[field]); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return r
}

func (c *CsvRecord) GetString(field string) string {
	data, ok := c.Record[field]
	if ok {
		return data
	} else {
		fmt.Println(field)
		return ""
	}
}

func LoadCsvCfg(filename string, row int) *CsvTable {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if reader == nil {
		fmt.Println(file)
		return nil
	}
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(records) < row {
		fmt.Println(filename, " is empty")
		return nil
	}
	colNum := len(records[0])
	recordNum := len(records)
	var allRecords []CsvRecord
	for i := row; i < recordNum; i++ {
		record := &CsvRecord{make(map[string]string)}
		for k := 0; k < colNum; k++ {
			record.Record[records[0][k]] = records[i][k]
		}
		allRecords = append(allRecords, *record)
	}
	var result = &CsvTable{
		filename,
		allRecords,
	}
	return result
}