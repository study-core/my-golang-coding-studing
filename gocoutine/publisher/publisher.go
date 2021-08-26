package publisher

import (
	"fmt"
	"my-golang-coding-studing/gocoutine/types"
	"sync"
	"time"
)

type Publisher struct {

	taskCache  map[string]*types.Task
	receiveTaskCh  <-chan *types.Task
	taskResultCh   chan *types.TaskResult
	reschs map[string]chan <- *types.TaskResult

	reschsLock sync.Mutex
}


func NewPublisher(receiveTaskCh chan *types.Task) *Publisher {

	pub := &Publisher{
		taskCache: make(map[string]*types.Task, 0),
		receiveTaskCh: receiveTaskCh,
		taskResultCh: make(chan *types.TaskResult, 100),
		reschs: make(map[string]chan <- *types.TaskResult, 0),
	}

	go pub.loop()

	return pub
}

func (p *Publisher) loop() {

	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case task := <- p.receiveTaskCh:
			p.handleTask(task, task.ResultCh)
		case res := <-p.taskResultCh:
			if nil == res {
				return
			}
			p.sendConsensusTaskResultToSched(res)
		case <-ticker.C:
			p.refresh()
		}
	}
}




func (p *Publisher) handleTask (task *types.Task, resCh chan <- *types.TaskResult) {

	p.taskCache[task.TaskId] = task

	p.addResCh(task.TaskId, resCh)

	fmt.Println("nima  打印到这里了吧 ???")


}

func (p *Publisher) addResCh( taskId string, resCh chan <- *types.TaskResult) {
	p.reschsLock.Lock()
	fmt.Println("AddTaskResultCh taskId: ", taskId)
	p.reschs[taskId] = resCh
	p.reschsLock.Unlock()
}


func (p *Publisher) collectTaskResultWillSendToSched(result *types.TaskResult) {
	p.taskResultCh <- result
}
func  (p *Publisher) sendConsensusTaskResultToSched (result *types.TaskResult) {
	p.reschsLock.Lock()
	fmt.Printf("Need SendTaskResultCh taskId: {%s}, result: {%s}\n", result.TaskId, result.String())
	if ch, ok := p.reschs[result.TaskId]; ok {
		fmt.Printf("Start SendTaskResultCh taskId: {%s}, result: {%s}\n", result.TaskId, result.String())
		ch <- result
		close(ch)
		delete( p.reschs, result.TaskId)
	}
	p.reschsLock.Unlock()
}


func (p *Publisher)  refresh () {

	for taskId, _ := range p.taskCache {

		res := &types.TaskResult{
			TaskId: taskId,
			Done: true,
		}
		delete(p.taskCache, taskId)
		p.collectTaskResultWillSendToSched(res)
	}

}