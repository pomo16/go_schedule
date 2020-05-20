package main

import (
	"go_schedule/consts"
	"go_schedule/crons"
)

func init() {
	crons.InitCrons(consts.CronType)
}

func main() {}
