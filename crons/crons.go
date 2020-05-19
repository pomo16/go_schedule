package crons

import (
	"github.com/robfig/cron"
	"go_schedule/consts"
	"go_schedule/process"
	"time"
)

//InitCrons 初始化定时任务
//TickType和LiteTickType效果一致，都基于time.Sleep，如果task执行时间超过定时间隔，定时执行逻辑会阻塞至任务完成并在任务完成后马上执行下一个任务，即在该情况下定时逻辑无效。
//CronType在上述情况下不会被阻塞，新任务会在新的goroutine中执行，不影响尚未执行完的上一个任务。
func InitCrons(cron int) {
	switch cron {
		case consts.CronType:
			CronType()
		case consts.TickType:
			TickType()
		case consts.LiteTickType:
			LiteTickType()
		default:
			//todo: log error
	}
}

//CronType 使用第三方cron包实现定时任务
func CronType() {
	//定时循环
	c := cron.New()
	spec := consts.CronTiming
	err := c.AddFunc(spec, taskFunc)
	if err != nil {
		panic(err)
	}
	c.Start()

	select {}
}

//TickType 使用内置time.Tick实现定时任务
func TickType() {
	tick := time.Tick(consts.TickTiming)
	for {
		select {
		case <-tick:
			taskFunc()
		}
	}

}

//LiteTickType 简化版time.Tick实现定时任务
func LiteTickType() {
	for range time.Tick(consts.TickTiming) {
		taskFunc()
	}
}

//taskFunc 要执行的任务函数
func taskFunc() {
	process.Schedule()
}


