// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package comm

import (
	"github.com/youhei-yp/wing/logger"
	"time"
)

// Task the type of task
type Task struct {
	queue     *Queue
	interrupt bool
	interval  time.Duration
	executing bool
}

var chexe = make(chan string)
var chnxt = make(chan string)

// TaskCallback task callback
type TaskCallback func(data interface{}) error

// GenTask generat a new task instance, you can set the interval duration
// and interrupt flag as the follow format:
// [CODE:]
//   interrupt := 1  // interrupt to execut the remain tasks when case error
//   interval := 500 // sleep interval between tasks in microseconds
//   task := comm.GenTask(callback, interrupt, interval)
//   task.Post(taskdata)
// [CODE]
func GenTask(callback TaskCallback, configs ...int) *Task {
	// generat the task and fill default configs
	task := &Task{
		queue: GenQueue(), interrupt: false, interval: 0, executing: false,
	}

	// set task configs from given data
	if configs != nil {
		task.interrupt = len(configs) > 0 && configs[0] > 0
		if len(configs) > 1 && configs[1] > 0 {
			task.interval = time.Duration(configs[1] * 1000)
		}
	}

	// start task channel to listen
	go task.innerTaskExecuter(callback)
	logger.I("Generat a task:{interrupt:", task.interrupt, ", interval:", task.interval, "}")
	return task
}

// Post post a task to tasks queue back
func (t *Task) Post(taskdata interface{}) {
	if taskdata == nil {
		return
	}
	t.queue.Push(taskdata)
	chexe <- "Post Action"
}

// SetInterrupt set interrupt flag
func (t *Task) SetInterrupt(interrupt bool) {
	t.interrupt = interrupt
}

// setInterval set wait interval between tasks in microseconds, and it must > 0.
func (t *Task) SetInterval(interval int) {
	if interval > 0 {
		t.interval = time.Duration(interval * 1000)
	}
}

// innerTaskExecuter task execte monitor to listen tasks
func (t *Task) innerTaskExecuter(callback TaskCallback) {
	for {
		logger.I("Blocking for select...")
		select {
		case action := <-chexe:
			logger.I("Received request from:", action)
			if callback == nil {
				logger.E("Nil task callback, abort request")
				break
			}

			// check current if executing status
			if t.executing {
				logger.W("Bussying now, try the next time...")
				break
			}

			// flag on executing and popup the topmost task to execte
			t.executing = true
			taskdata, err := t.queue.Pop()
			if err != nil {
				t.executing = false
				logger.I("Executed all task")
				break
			}

			if err := callback(taskdata); err != nil {
				logger.E("Execute task callback err:", err)
				if t.interrupt {
					logger.I("Interrupted tasks on error")
					t.executing = false
					break
				}
			}
			if t.interval > 0 {
				logger.I("Waiting to execute the next task after:", t.interval)
				time.Sleep(t.interval)
			}
			t.executing = false
			chnxt <- "Post Next"

		case <-chnxt:
			logger.I("Received next require, post action...")
			chexe <- "Next Action"
		}
		logger.I("Exit task select!")
	}
}
