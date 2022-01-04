package cron

import (
	"github.com/go-co-op/gocron"
	"go_scaffold/app/cron/cron_job"
	"log"
	"time"
)

//	@method Kernel
//	@description: cron kernel
//	@return error
func Kernel() error {

	// 实例化一个定时任务调度器
	sch := gocron.NewScheduler(time.Local)

	// 让全部任务在项目启动时不运行
	sch.WaitForScheduleAll()

	sch.Every(5).Seconds().Do(cron_job.Every5Seconds)

	log.Print("cron server is running ")

	// 作为独立的服务运行，那么使用阻塞模式
	sch.StartBlocking()

	// 如果是整合到其他服务运行，例如和http服务一起运行那么使用异步模式启动即可
	//sch.StartAsync()

	return nil
}
