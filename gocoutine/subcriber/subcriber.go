package subcriber

import (
	"fmt"
	"my-golang-coding-studing/gocoutine/types"
)

type Subcriber struct {
	SendTaskCh     chan *types.Task
}

func NewSubcriber (sendTaskCh chan *types.Task) *Subcriber {
	return &Subcriber{
		SendTaskCh: sendTaskCh,
	}
}

func (s *Subcriber) SendTask() {

	task := &types.Task{
		TaskId: "Gavin love kally task",
		ResultCh: make(chan *types.TaskResult, 0),
	}
	s.sendTask(task)
	res := task.ReceiveResult()

	fmt.Println("Received the res:", res.TaskId, res.Done)
}

func (s *Subcriber) sendTask (task *types.Task) {
	s.SendTaskCh <- task
}
