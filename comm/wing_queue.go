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
	"container/list"
	"fmt"
	"github.com/youhei-yp/wing/invar"
	"sync"
)

// Queue the type of queue with sync lock
type Queue struct {
	list  *list.List
	mutex sync.Mutex
}

// GenQueue generat a new queue instance
func GenQueue() *Queue {
	return &Queue{
		list: list.New(),
	}
}

// Push push a data to queue back if the data not nil
func (q *Queue) Push(data interface{}) {
	if data == nil {
		return
	}
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.list.PushBack(data)
}

// Pop pick and remove the front data of queue,
// it will return invar.ErrEmptyData error if the queue is empty
func (q *Queue) Pop() (interface{}, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if e := q.list.Front(); e != nil {
		q.list.Remove(e)
		return e.Value, nil
	}
	return nil, invar.ErrEmptyData
}

// Pick pick but not remove the front data of queue,
// it will return invar.ErrEmptyData error if the queue is empty
func (q *Queue) Pick() (interface{}, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if e := q.list.Front(); e != nil {
		return e.Value, nil
	}
	return nil, invar.ErrEmptyData
}

// Clear clear the queue all data
func (q *Queue) Clear() {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for e := q.list.Front(); e != nil; {
		en := e.Next()
		q.list.Remove(e)
		e = en
	}
}

// Len return the length of queue
func (q *Queue) Len() int {
	return q.list.Len()
}

// Dump print out the queue data.
// this method maybe just use for debug to out put queue items
func (q *Queue) Dump() {
	fmt.Println("-- dump the queue: (front -> back)")
	for e := q.list.Front(); e != nil; e = e.Next() {
		fmt.Println(fmt.Sprintf("   : %v", e.Value))
	}
}
