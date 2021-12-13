package game

import (
	"fmt"
	"server/csvs"
)

type Weapon struct {
	WeaponId	int
	KeyId	int
}

type ModWeapon struct {
	WeaponInfo map[int]*Weapon
	MaxKey int
}

func (self *ModWeapon) AddItem(itemId int, num int64)  {
	config:=csvs.GetWeaponConfig(itemId)
	if config==nil{
		fmt.Println("配置不存在")
		return
	}

	if len(self.WeaponInfo)+int(num) > csvs.WEAPON_MAX_COUNT {
		fmt.Println("超过最大值")
		return
	}

	for i := int64(0); i < num; i++ {
		weapon := new(Weapon)
		weapon.WeaponId = itemId
		self.MaxKey++
		weapon.KeyId = self.MaxKey
		self.WeaponInfo[weapon.KeyId] = weapon
		fmt.Println("获得武器:", csvs.GetItemName(itemId), "------武器编号:", weapon.KeyId)
	}
}