package queue

import "sync"

//SliceQueue 切片队列
type SliceQueue struct {
	data  []interface{}
	mutex *sync.Mutex
}

//NewSliceQueue 初始化队列
func NewSliceQueue(len int, data ...interface{}) *SliceQueue {
	queue := &SliceQueue{
		data:  make([]interface{}, 0, len),
		mutex: &sync.Mutex{},
	}
	for _, v := range data {
		queue.data = append(queue.data, v)
	}
	return queue
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

//Pop 移除队列顶元素
func (s *SliceQueue) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	item := s.data[0]
	s.data = s.data[1:]
	return item
}

//BFS 广度优先搜索
func (s *SliceQueue) BFS(fn func(front interface{})) {
	for !s.Empty() {
		front := s.Pop()
		if fn != nil {
			fn(front)
		}
	}
}
