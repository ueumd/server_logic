package csvs

import "server/utils"

/**
武器
weapon
n. 武器, 兵器
 */
type ConfigWeapon struct {
	WeaponId int `json:"weaponId"`
	Type int `json:"type"`
	Star int `json:"star"`
}

var (
	ConfigWeaponMap map[int]*ConfigWeapon
)

func init()  {
	ConfigWeaponMap = make(map[int]*ConfigWeapon)
	utils.GetCsvUtilMgr().LoadCsv("Weapon", &ConfigWeaponMap)
	return
}

func GetWeaponConfig(weaponId int) *ConfigWeapon  {
	return ConfigWeaponMap[weaponId]
}