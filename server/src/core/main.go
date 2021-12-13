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
	// 协程更新违禁词库
	go game.GetManageBanWord().Run()

	fmt.Println("数据测试-----start\n")

	playerTest := game.NewTestPlayer()
	go playerTest.Run()
	for {

	}
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
