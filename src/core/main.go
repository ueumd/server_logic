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
	// 3 签名      			------- 多线程(5) banword
	// 4 名字
	// 5 冒险等级 冒险阅历	-------	map(10)
	// 6 世界等级 冷却时间
	// 7 生日
	// 8 展示阵容 展示名片

	// 原神 -- 成长核心：背包系统

	// 业务算法：公主连接 夏日赛跑

	// *******************************************

	// 当前模块：背包
	// 1. 物品识别
	// 2. 物品增加
	// 2. 物品消耗
	// 2. 物品使用
	// 2. 角色模块 -> 头像模块

	// 加载配置
	// 执行csv所有init
	csvs.CheckLoadCsv()

	fmt.Println("数据测试-----start\n")

	// 协程更新违禁词库
	go game.GetManageBanWord().Run()

	playerGM := game.NewTestPlayer()

	// 测试名片
	//playerGM.ModPlayer.SetShowCard([]int{1001, 1001, 1001, 1002, 1001, 1005}, playerGM)
	//playerGM.ModPlayer.SetShowCard([]int{}, playerGM)
	//playerGM.ModPlayer.SetShowCard([]int{1009}, playerGM)

	// 测试阵容
	// playerGM.ModPlayer.SetShowTeam([]int{1001, 1001, 1001, 1002, 1001, 1005, 1001, 1001, 1002, 10051001, 1001, 1002, 1005}, playerGM)
	// playerGM.ModPlayer.SetShowTeam([]int{1001, 1001, 1001, 1002}, playerGM)

	// 测试生日
	//playerGM.ModPlayer.SetBirth(2000)
	//playerGM.ModPlayer.SetBirth(1235)
	//playerGM.ModPlayer.SetBirth(10)
	//playerGM.ModPlayer.SetBirth(1126)
	//playerGM.ModPlayer.SetBirth(520)

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

	// 玩家等级测试
	// 直接升级到 60级
	//playerGM.ModPlayer.AddExp(10000000, playerGM)

	// 玩家等级测试
	// ticker := time.NewTicker(time.Second * 3)
	//for {
	//	select {
	//	case <-ticker.C:
	//		playerGM.ModPlayer.AddExp(5000)
	//	}
	//}

	// 突破任务测试 map读写问题
	// go playerSet(playerGM)
	// go playerGet(playerGM)

	// go playerSetLock(playerGM)
	// go playerGetLock(playerGM)

	// config读取不加锁
	//go playerLoadConfig(playerGM)
	//go playerLoadConfig(playerGM)
	//go playerLoadConfig(playerGM)
	//go playerLoadConfig(playerGM)
	//go playerLoadConfig(playerGM)

	// 玩家世界等级测试
	//ticker := time.NewTicker(time.Second * 1)
	//for {
	//	select {
	//	case <-ticker.C:
	//		if time.Now().Unix()%3 == 0 {
	//			// 降
	//			playerGM.ReturnWorldLevel()
	//
	//		} else if time.Now().Unix()%5 == 0 {
	//			playerGM.ReduceWorldLevel()
	//		}
	//	}
	//}

	// 玩家登录
	// 每10S 加入一个玩家
	//ticker := time.NewTicker(time.Second * 10)
	//for {
	//	select {
	//	case <-ticker.C:
	//		playerTest := game.NewTestPlayer()
	//		fmt.Println("==================== Player ====================")
	//		// 玩家启动自己的协程
	//		go playerTest.Run()
	//	}
	//}

	// 背包：物品添加
	playerGM.ModBag.AddItem(1000001, playerGM)
	playerGM.ModBag.AddItem(1000006, playerGM)
	playerGM.ModBag.AddItem(1000008, playerGM)
	playerGM.ModBag.AddItem(2000002, playerGM)
	playerGM.ModBag.AddItem(2000021, playerGM)
	playerGM.ModBag.AddItem(3000001, playerGM)
	playerGM.ModBag.AddItem(3000002, playerGM)
	playerGM.ModBag.AddItem(3000003, playerGM)
	playerGM.ModBag.AddItem(3000044, playerGM)
	playerGM.ModBag.AddItem(4000001, playerGM)
	playerGM.ModBag.AddItem(4000002, playerGM)

	// 确实协程在不断运行
	//for {
	//	//
	//}

	return

}

// ********************测试Map读写************************

// 突破任务测试 map并发读写安全问题
// fatal error: concurrent map read and map write

func playerSet(player *game.Player) {
	for i := 0; i < 1000000; i++ {
		player.ModUniqueTask.MyTaskInfo[10001] = new(game.TaskInfo)
	}
}

func playerGet(player *game.Player) {
	for i := 0; i < 1000000; i++ {
		_, ok := player.ModUniqueTask.MyTaskInfo[10001]
		if ok {
		}
	}
}

// 加上 读写锁 有性能消耗
func playerSetLock(player *game.Player) {
	startTime := time.Now().Nanosecond()
	for i := 0; i < 1000000; i++ {
		// 加上锁
		player.ModUniqueTask.Locker.Lock()
		player.ModUniqueTask.MyTaskInfo[10001] = new(game.TaskInfo)
		player.ModUniqueTask.Locker.Unlock()
	}
	endTime := time.Now().Nanosecond() - startTime

	fmt.Println(endTime / 1000000)

}

func playerGetLock(player *game.Player) {
	startTime := time.Now().Nanosecond()
	for i := 0; i < 1000000; i++ {
		player.ModUniqueTask.Locker.RLock()
		_, ok := player.ModUniqueTask.MyTaskInfo[10001]
		if ok {
			//
		}
		player.ModUniqueTask.Locker.RUnlock()
	}
	endTime := time.Now().Nanosecond() - startTime

	fmt.Println(endTime / 1000000)
}

// 对于配置模块所有玩家都是在读 不加锁
func playerLoadConfig(player *game.Player) {
	for i := 0; i < 1000000; i++ {
		config := csvs.ConfigUniqueTaskMap[10001]

		if config != nil {
			fmt.Println(config.TaskId)
		}
	}
}
