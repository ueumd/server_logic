package csvs

import "server/utils"

type ConfigIcon struct {
	IconId int `json:"IconId"`
	Check 	int 	`json:"check"`
}

var (
	ConfigIconMap map[int]*ConfigIcon
	ConfigIconMapByRoleId map[int]*ConfigIcon
)

func GetIconConfig(iconId int) *ConfigIcon {
	return ConfigIconMap[iconId]
}

func GetIconConfigByRoleId(roleId int) *ConfigIcon {
	return ConfigIconMapByRoleId[roleId]
}

//func loadIconCsv() {
//	exPath, _ := os.Getwd()
//	fmt.Println("path", exPath)
//
//	fs, _ := os.Open("../../csv/Icon.csv")
//	result := csv.NewReader(fs)
//
//	content, err := result.ReadAll()
//	if err != nil {
//		log.Fatalf("can not readall, err is %+v", err)
//	}
//	ConfigIconMap = make(map[int]*ConfigIcon)
//	for index, row := range content {
//		if index > 2 {
//			IconId, _ := strconv.Atoi(row[0])
//			ConfigIconMap[IconId] = &ConfigIcon{
//				IconId,
//			}
//
//		}
//	}
//
//	fmt.Println(ConfigIconMap[0])
//}

func init() {
	ConfigIconMap = make(map[int]*ConfigIcon)

	utils.GetCsvUtilMgr().LoadCsv("Icon", &ConfigIconMap)

	ConfigIconMapByRoleId = make(map[int]*ConfigIcon)

	for _, v := range ConfigIconMap {
		ConfigIconMapByRoleId[v.Check] = v
	}

	return
}
