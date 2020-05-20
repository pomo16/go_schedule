package process

import (
	"context"
	"go_schedule/exceptions"
	"go_schedule/model"
)

//Load 任务灌注
func Load(ctx context.Context, done <-chan struct{}) (<-chan *model.TaskS, <-chan error) {
	taskc := make(chan *model.TaskS)
	errc := make(chan error, 1)

	go func() {
		defer close(taskc)
		errc <- func(ctx context.Context) error {
			//Add tasks to the channel
			taskList := []string{"task1", "task2", "task3"}
			for _, v := range taskList {
				select {
				case taskc <- &model.TaskS{Task: v}:
				case <-done:
					return exceptions.ErrTaskAbort
				}
			}
			return nil
		}(ctx)
	}()

	return taskc, errc
}
