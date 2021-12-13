package game

import (
	"fmt"
	"server/csvs"
)

type RoleInfo struct {
	RoleId	int
	GetTimes int
}

// 角色
type ModeRole struct {
	RoleInfo map[int]*RoleInfo
}

func (self *ModeRole) IsHasRole(roleId int) bool {
	return true
}

// 角色等级
func (self *ModeRole) GetRoleLevel(roleId int) int {
	return 80
}


func (self *ModeRole) AddItem(roleId int, num int64) {
	config := csvs.GetRoleConfig(roleId)
	if config == nil {
		fmt.Println("配置不存在roleId:", roleId)
		return
	}
	for i:=0; i<int(num); i++ {
		_, ok := self.RoleInfo[roleId]

		if !ok {
			// 第一次
			data := new(RoleInfo)
			data.RoleId = roleId
			data.GetTimes = 1
			self.RoleInfo[roleId] = data
		} else {
			fmt.Println("获得实际物品")
			self.RoleInfo[roleId].GetTimes ++
		}
	}

	itemConfig := csvs.GetItemConfig(roleId)
	if itemConfig != nil {
		fmt.Println("获得角色", itemConfig.ItemName, "次数", roleId, self.RoleInfo[roleId].GetTimes, "次")
	}
}
