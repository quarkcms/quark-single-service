package scheduler

import (
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	Cron *gocron.Scheduler
}

var once sync.Once
var scheduler *Scheduler

// 初始化调度器
func NewScheduler() *Scheduler {
	// 单例模式初始化调度器
	once.Do(func() {
		// 设置时区
		location, _ := time.LoadLocation("Asia/Shanghai")
		scheduler = &Scheduler{
			Cron: gocron.NewScheduler(location),
		}
	})

	return scheduler
}

// 启动调度器
func (p *Scheduler) Start() {
	p.Cron.StartAsync()
}
