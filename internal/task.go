package internal

import (
	"github.com/robfig/cron"
	"sync"
)

type TaskFunc func(params ...interface{})

var (
	taskList chan *TaskExecutor
	once     sync.Once
	onceCron sync.Once
	taskCron *cron.Cron
)

//TaskExecutor 任务组件
type TaskExecutor struct {
	f        TaskFunc
	p        []interface{}
	callback func()
}

func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}

func init() {
	ch := getTaskList()
	go func() {
		for t := range ch {
			doTask(t)
		}
	}()
}

func doTask(t *TaskExecutor) {
	go func() {
		defer func() {
			if t.callback != nil {
				t.callback()
			}
		}()
		t.Exec()
	}()
}
func NewTaskExecutor(f TaskFunc, p []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{f: f, p: p, callback: callback}
}

func (t *TaskExecutor) Exec() {
	t.f(t.p...)
}

func Task(f TaskFunc, callback func(), params ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		getTaskList() <- NewTaskExecutor(f, params, callback)
	}()
}

func getCronTask() *cron.Cron {
	onceCron.Do(func() {
		taskCron = cron.New()
	})
	return taskCron
}
