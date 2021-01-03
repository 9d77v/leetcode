package queue

//SliceQueue 队列
type SliceQueue struct {
	data []interface{}
}

//NewSliceQueue 初始化队列
func NewSliceQueue(len int) *SliceQueue {
	return &SliceQueue{
		data: make([]interface{}, 0, len),
	}
}

//Len 长度
func (s *SliceQueue) Len() int {
	return len(s.data)
}

//Empty 判断队列是否为空
func (s *SliceQueue) Empty() bool {
	return s.Len() == 0
}

//Push 新增元素
func (s *SliceQueue) Push(v interface{}) {
	s.data = append(s.data, v)
}

//Pop 移除队列头元素
func (s *SliceQueue) Pop() interface{} {
	if s.Len() == 0 {
		return nil
	}
	v := s.data[0]
	if s.Len() == 1 {
		s.data = s.data[0:0]
	} else {
		s.data = s.data[1:s.Len()]
	}
	return v
}

//Peek 获取队列头元素
func (s *SliceQueue) Peek() interface{} {
	if s.Len() == 0 {
		return nil
	}
	return s.data[0]
}

//Clear 清空队列
func (s *SliceQueue) Clear() {
	s.data = s.data[0:0]
}
