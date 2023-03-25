package log_

import (
	"fmt"
	"log"
)

type SubTask struct {
	*Task
}
type Task struct {
	TaskId int
	*log.Logger
}

func (task *SubTask) Printf(format string, args ...interface{}) {
	task.Logger.Printf("%v %v", task.TaskId, fmt.Sprintf(format, args...))
}
