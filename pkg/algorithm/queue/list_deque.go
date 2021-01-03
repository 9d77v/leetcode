package queue

import "container/list"

//ListDeque list实现的双端队列
type ListDeque struct {
	data *list.List
	cap  int
}

//NewListDeque 初始化队列
func NewListDeque(cap int) *ListDeque {
	return &ListDeque{
		data: list.New(),
		cap:  cap,
	}
}

//Cap 容量
func (s *ListDeque) Cap() int {
	return s.cap
}

//Len 长度
func (s *ListDeque) Len() int {
	return s.data.Len()
}

//Empty 判断队列是否为空
func (s *ListDeque) Empty() bool {
	return s.Len() == 0
}

//PushFront 新增元素
func (s *ListDeque) PushFront(v interface{}) {
	s.data.PushFront(v)
}

//PushBack 新增元素
func (s *ListDeque) PushBack(v interface{}) {
	s.data.PushBack(v)
}

//PopFront 移除队头元素
func (s *ListDeque) PopFront() interface{} {
	if s.Empty() {
		return nil
	}
	v := s.data.Front()
	s.data.Remove(v)
	return v.Value
}

//PopBack 移除队尾元素
func (s *ListDeque) PopBack() interface{} {
	if s.Empty() {
		return nil
	}
	v := s.data.Back()
	s.data.Remove(v)
	return v.Value
}

//Front 返回队头元素
func (s *ListDeque) Front() interface{} {
	return s.data.Front().Value
}

//Back 返回队尾元素
func (s *ListDeque) Back() interface{} {
	return s.data.Back().Value
}
