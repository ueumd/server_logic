package main

import "fmt"
import "server_logic/game"

func main() {
	// 当前模块基础信息
	// 1 UID
	// 2 头像 名片
	// 3 签名
	// 4 名字
	// 5 冒险等级 冒险阅历
	// 6 世界等级 冷却时间
	// 7 生日
	// 8 展示阵容 展示名片

	fmt.Println("数据测试-----start\n")

	player := game.NewTestPlayer()

	//player.RecvSetIcon(1) // 胡桃
	//player.RecvSetIcon(2) //
	//player.RecvSetIcon(3)
	//
	//player.RecvSetCard(1)
	//player.RecvSetCard(2)
	//player.RecvSetCard(3)

	// 改名
	player.RecvSetName("好人")
	player.RecvSetName("坏蛋")
	player.RecvSetName("求外挂")
	player.RecvSetName("好玩")
	player.RecvSetName("TMD")

}
