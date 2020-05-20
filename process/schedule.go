package process

import (
	"context"
	"github.com/sirupsen/logrus"
	"go_schedule/consts"
	"go_schedule/model"
	"sync"
)

//Schedule 任务调度入口
func Schedule() {
	ctx := context.Background()
	donec := make(chan struct{})
	defer close(donec)

	taskc, errc := Load(ctx, donec)
	resultc := make(chan *model.ResultS)

	var wg sync.WaitGroup
	wg.Add(consts.WorkerNum)
	for i := 0; i < consts.WorkerNum; i++ {
		go func() {
			process(donec, taskc, resultc)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		//关闭之前，读结果的for...range...会一直阻塞
		close(resultc)
	}()

	//parse result
	for r := range resultc {
		logrus.Infoln(r)
	}

	if err := <-errc; err != nil {
		logrus.Errorf("schedule process error[%v]", err)
	}
}

func process(donec <-chan struct{}, taskc <-chan *model.TaskS, resultc chan<- *model.ResultS) {
	for task := range taskc {
		select {
		case resultc <- doTask(task):
		case <-donec:
			return
		}
	}
}

func doTask(task *model.TaskS) *model.ResultS {
	result := &model.ResultS{Result: task.Task + " done"}
	return result
}
