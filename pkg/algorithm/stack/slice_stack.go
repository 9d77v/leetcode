package stack

import "sync"

//SliceStack 切片栈
type SliceStack struct {
	data  []interface{}
	mutex *sync.Mutex
}

//NewSliceStack 初始化栈
func NewSliceStack(len int) *SliceStack {
	return &SliceStack{
		data:  make([]interface{}, 0, len),
		mutex: &sync.Mutex{},
	}
}

//Len 长度
func (s *SliceStack) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.data)
}

//IsEmpty 判断栈是否为空
func (s *SliceStack) IsEmpty() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.data) == 0
}

//Push 新增元素
func (s *SliceStack) Push(v interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = append(s.data, v)
}

//Pop 移除栈顶元素
func (s *SliceStack) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

//Peek 获取栈顶元素
func (s *SliceStack) Peek() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

//Clear 清空栈
func (s *SliceStack) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = s.data[0:0]
}
