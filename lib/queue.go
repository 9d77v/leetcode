package lib

//Queue 队列
type Queue struct {
	data []interface{}
}

//NewQueue 初始化队列
func NewQueue() *Queue {
	return &Queue{
		data: make([]interface{}, 0),
	}
}

//Len 获取队列长度
func (s *Queue) Len() int {
	return len(s.data)
}

//Push 新增元素
func (s *Queue) Push(v interface{}) {
	s.data = append(s.data, v)
}

//Pop 移除队列顶元素
func (s *Queue) Pop() interface{} {
	size := len(s.data)
	item := s.data[0]
	if size == 1 {
		s.data = s.data[0:0]
	}
	s.data = s.data[1:size]
	return item
}
