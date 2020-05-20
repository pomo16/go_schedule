package consts

import "time"

//Cron 定时任务种类
const (
	CronType = iota
	TickType
	LiteTickType
)

//Timing 定时时机
const (
	CronTiming = "0/5 * * * *" //5s一次
	TickTiming = 5 * time.Second
)

//TimeStr 时间格式化模板
const TimeStr = "2006-01-02 15:04:05"

//log文件相关
const (
	LogFilePath    = "go_schedule_log"
	LogFilePrefix  = "schedule"
	LogFileTimeStr = "20060102150405"
	LogFileSuffix  = ".log"
)

//任务调度参数
const (
	WorkerNum = 10
)
