package game

// 觉色
type ModeRole struct {
}

func (self *ModeRole) IsHasRole(roleId int) bool {
	return true
}

// 觉色等级
func (self *ModeRole) GetRoleLevel(roleId int) int {
	return 80
}
