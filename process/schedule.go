package process

import (
	"fmt"
	"go_schedule/consts"
	"time"
)

//Schedule 任务调度入口
func Schedule() {
	fmt.Println(time.Now().Format(consts.TimeStr))
}
