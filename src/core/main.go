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


	// 玩家等级测试
	// 直接升级到 60级
	playerGM.ModPlayer.AddExp(10000000, playerGM)

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

	go playerLoadConfig(playerGM)
	go playerLoadConfig(playerGM)
	go playerLoadConfig(playerGM)
	go playerLoadConfig(playerGM)
	go playerLoadConfig(playerGM)

	// 确实协程在不断运行
	for {
		//
	}

	return

}



// ********************测试************************

// 突破任务测试 map并发读写安全问题
// fatal error: concurrent map read and map write

func playerSet(player *game.Player)  {
	for i:=0; i<1000000; i++ {
		player.ModUniqueTask.MyTaskInfo[10001] = new(game.TaskInfo)
	}
}

func playerGet(player *game.Player)  {
	for i:=0; i<1000000; i++ {
		_, ok := player.ModUniqueTask.MyTaskInfo[10001]
		if ok {
		}
	}
}

// 加上 读写锁 有性能消耗
func playerSetLock(player *game.Player)  {
	startTime := time.Now().Nanosecond()
	for i:=0; i<1000000; i++ {
		// 加上锁
		player.ModUniqueTask.Locker.Lock()
		player.ModUniqueTask.MyTaskInfo[10001] = new(game.TaskInfo)
		player.ModUniqueTask.Locker.Unlock()
	}
	endTime := time.Now().Nanosecond() - startTime

	fmt.Println(endTime/1000000)

}

func playerGetLock(player *game.Player)  {
	startTime := time.Now().Nanosecond()
	for i:=0; i<1000000; i++ {
		player.ModUniqueTask.Locker.RLock()
		_, ok := player.ModUniqueTask.MyTaskInfo[10001]
		if ok {
			//
		}
		player.ModUniqueTask.Locker.RUnlock()
	}
	endTime := time.Now().Nanosecond() - startTime

	fmt.Println(endTime/1000000)
}


// 对于配置模块所有玩家都是在读 不加锁
func playerLoadConfig(player *game.Player)  {
	for i:=0; i<1000000; i++ {
		config := csvs.ConfigUniqueTaskMap[10001]

		if config != nil{
			fmt.Println(config.TaskId)
		}
	}
}