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

// Stack the type of stack with sync lock
type Stack struct {
	list  *list.List
	mutex sync.Mutex
}

// GenStack generat a new stack instance
func GenStack() *Stack {
	return &Stack{
		list: list.New(),
	}
}

// Push push a data to stack top one if the data not nil
func (s *Stack) Push(data interface{}) {
	if data == nil {
		return
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.list.PushFront(data)
}

// Pop pick and remove the top data of stack,
// it will return invar.ErrEmptyData error if the stack is empty
func (s *Stack) Pop() (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if e := s.list.Front(); e != nil {
		s.list.Remove(e)
		return e.Value, nil
	}
	return nil, invar.ErrEmptyData
}

// Pick pick but not remove the top data of stack,
// it will return invar.ErrEmptyData error if the stack is empty
func (s *Stack) Pick() (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if e := s.list.Front(); e != nil {
		return e.Value, nil
	}
	return nil, invar.ErrEmptyData
}

// Clear clear the stack all data
func (s *Stack) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for e := s.list.Front(); e != nil; {
		en := e.Next()
		s.list.Remove(e)
		e = en
	}
}

// Len return the length of stack
func (s *Stack) Len() int {
	return s.list.Len()
}

// Dump print out the stack data.
// this method maybe just use for debug to out put stack items
func (s *Stack) Dump() {
	fmt.Println("-- dump the stack: (top -> bottom)")
	for e := s.list.Front(); e != nil; e = e.Next() {
		fmt.Println(fmt.Sprintf("   : %v", e.Value))
	}
}
