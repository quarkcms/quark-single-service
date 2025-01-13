package task

import (
	"github.com/quarkcloudio/quark-smart/v2/config"
	"github.com/quarkcloudio/quark-smart/v2/pkg/scheduler"
)

// 注册任务
func RegisterTask() {
	// 初始化调度器
	s := scheduler.NewScheduler()

	if config.App.Pro {
		// 账单任务
		RunBillTask(s.Cron)
	}

	// 启动调度器
	s.Start()
}
