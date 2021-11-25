package main

import (
	"fmt"
	"server/csvs"
	"server/game"
	"time"
)

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

	// 加载配置

	fmt.Println("数据测试-----start\n")

	// 执行csv所有init
	csvs.CheckLoadCsv()

	// 协程更新词库
	go game.GetManageBanWord().Run()

	playerGM := game.NewTestPlayer()

	//playerGM.RecvSetIcon(1) // 胡桃
	//playerGM.RecvSetIcon(2) //
	//playerGM.RecvSetIcon(3)
	//
	//playerGM.RecvSetCard(1)
	//playerGM.RecvSetCard(2)
	//playerGM.RecvSetCard(3)

	// 改名
	//playerGM.RecvSetName("好人")
	//playerGM.RecvSetName("坏蛋")
	//playerGM.RecvSetName("求外挂")
	//playerGM.RecvSetName("好玩")
	//playerGM.RecvSetName("TMD")

	// 测试
	//tickerIn := time.NewTicker(time.Second * 3) // 3S
	//tickerOut := time.NewTicker(time.Second * 5)
	//
	//for {
	//	select {
	//	case <-tickerIn.C:
	//		playerGM.RecvSetIcon(int(time.Now().Unix()))
	//	case <-tickerOut.C:
	//		playerGM.RecvSetName("r u ok")
	//	}
	//}

	// 测试违禁词库
	//ticker := time.NewTicker(time.Second * 1)
	//for {
	//	select {
	//	case <-ticker.C:
	//		if time.Now().Unix()%3 == 0 {
	//			playerGM.RecvSetName("专业代练")
	//		} else if time.Now().Unix()%5 == 0 {
	//			playerGM.RecvSetName("良民")
	//		}
	//	}
	//}

	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-ticker.C:
			playerGM.ModPlayer.AddExp(5000)
		}
	}

}
