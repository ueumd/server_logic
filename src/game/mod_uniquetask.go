package game

import "sync"

// 记录任务状态
type TaskInfo struct {
	TaskId int
	State  int
}
type ModUniqueTask struct {
	MyTaskInfo map[int]*TaskInfo
	Locker *sync.RWMutex	// 读写锁 解决map 读写问题
}

func (self *ModUniqueTask) IsTaskFinish(taskId int) bool {

	// 测试代码 卡在45级左右
	if taskId == 10001 || taskId == 10002 {
		return  true
	}
	// 测试代码

	task, ok := self.MyTaskInfo[taskId]

	if !ok {
		return false
	}

	return task.State == TASK_STATE_FINISH
}
