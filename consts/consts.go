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


