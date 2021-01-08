package queue

import "sync"

//SliceDeque slice实现的双端队列
type SliceDeque struct {
	data  []interface{}
	cap   int
	mutex *sync.Mutex
}

//NewSliceDeque 初始化队列
func NewSliceDeque(cap int) *SliceDeque {
	return &SliceDeque{
		data:  make([]interface{}, 0, cap),
		cap:   cap,
		mutex: &sync.Mutex{},
	}
}

//Cap 容量
func (s *SliceDeque) Cap() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.cap
}

//Len 长度
func (s *SliceDeque) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.data)
}

//Empty 判断队列是否为空
func (s *SliceDeque) Empty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.data) == 0
}

//PushFront 新增元素
func (s *SliceDeque) PushFront(v interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = append([]interface{}{v}, s.data...)
}

//PushBack 新增元素
func (s *SliceDeque) PushBack(v interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = append(s.data, v)
}

//PopFront 移除队头元素
func (s *SliceDeque) PopFront() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	v := s.data[0]
	if len(s.data) == 1 {
		s.data = s.data[0:0]
	} else {
		s.data = s.data[1:]
	}
	return v
}

//PopBack 移除队尾元素
func (s *SliceDeque) PopBack() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

//Front 返回队头元素
func (s *SliceDeque) Front() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	return s.data[0]
}

//Back 返回队尾元素
func (s *SliceDeque) Back() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}
