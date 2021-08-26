package types

import "fmt"

type Task struct {
   TaskId  string
   ResultCh chan *TaskResult
}

func (t *Task) ReceiveResult () *TaskResult {
	return <- t.ResultCh
}

type TaskResult struct {
	TaskId  string
	Done    bool
}

func (tr *TaskResult) String() string {
	return fmt.Sprintf(`{"taskId": %s, "done": %v}`, tr.TaskId, tr.Done)
}