package queue

//SliceDeque slice实现的双端队列
type SliceDeque struct {
	data []interface{}
	cap  int
}

//NewSliceDeque 初始化队列
func NewSliceDeque(cap int) *SliceDeque {
	return &SliceDeque{
		data: make([]interface{}, 0, cap),
		cap:  cap,
	}
}

//Cap 容量
func (s *SliceDeque) Cap() int {
	return s.cap
}

//Len 长度
func (s *SliceDeque) Len() int {
	return len(s.data)
}

//Empty 判断队列是否为空
func (s *SliceDeque) Empty() bool {
	return s.Len() == 0
}

//PushFront 新增元素
func (s *SliceDeque) PushFront(v interface{}) {
	s.data = append([]interface{}{v}, s.data...)
}

//PushBack 新增元素
func (s *SliceDeque) PushBack(v interface{}) {
	s.data = append(s.data, v)
}

//PopFront 移除队头元素
func (s *SliceDeque) PopFront() interface{} {
	if s.Empty() {
		return nil
	}
	v := s.data[0]
	if s.Len() == 1 {
		s.data = s.data[0:0]
	} else {
		s.data = s.data[1:]
	}
	return v
}

//PopBack 移除队尾元素
func (s *SliceDeque) PopBack() interface{} {
	if s.Empty() {
		return nil
	}
	v := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]
	return v
}

//Front 返回队头元素
func (s *SliceDeque) Front() interface{} {
	if s.Empty() {
		return nil
	}
	return s.data[0]
}

//Back 返回队尾元素
func (s *SliceDeque) Back() interface{} {
	if s.Empty() {
		return nil
	}
	return s.data[s.Len()-1]
}
