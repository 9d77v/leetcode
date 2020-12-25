package lib

import (
	"errors"
	"sync"
)

//Queue 队列
type Queue struct {
	data  []interface{}
	mutex sync.Mutex
}

//NewQueue 初始化队列
func NewQueue() *Queue {
	return &Queue{
		data:  make([]interface{}, 0),
		mutex: sync.Mutex{},
	}
}

//Len 获取队列长度
func (s *Queue) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.data)
}

//Push 新增元素
func (s *Queue) Push(v interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = append(s.data, v)
}

//Pop 移除队列顶元素
func (s *Queue) Pop() (item interface{}, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	size := len(s.data)
	if size == 0 {
		err = errors.New("stack is empty")
		return
	}
	item = s.data[0]
	if size == 1 {
		s.data = s.data[0:0]
	}
	s.data = s.data[1:size]
	return item, nil
}
