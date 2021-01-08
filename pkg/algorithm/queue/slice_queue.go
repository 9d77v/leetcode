package queue

import "sync"

//SliceQueue 队列
type SliceQueue struct {
	data  []interface{}
	mutex *sync.Mutex
}

//NewSliceQueue 初始化队列
func NewSliceQueue(len int) *SliceQueue {
	return &SliceQueue{
		data:  make([]interface{}, 0, len),
		mutex: &sync.Mutex{},
	}
}

//Len 长度
func (s *SliceQueue) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.data)
}

//Empty 判断队列是否为空
func (s *SliceQueue) Empty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.data) == 0
}

//Push 新增元素
func (s *SliceQueue) Push(v interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = append(s.data, v)
}

//Pop 移除队列头元素
func (s *SliceQueue) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	v := s.data[0]
	if len(s.data) == 1 {
		s.data = s.data[0:0]
	} else {
		s.data = s.data[1:len(s.data)]
	}
	return v
}

//Peek 获取队列头元素
func (s *SliceQueue) Peek() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	return s.data[0]
}

//Clear 清空队列
func (s *SliceQueue) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = s.data[0:0]
}
