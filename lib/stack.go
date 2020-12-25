package lib

//Stack 栈
type Stack struct {
	data []interface{}
}

//NewStack 初始化栈
func NewStack(len int) *Stack {
	return &Stack{
		data: make([]interface{}, 0, len),
	}
}

//IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

//Push 新增元素
func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

//Pop 移除栈顶元素
func (s *Stack) Pop() interface{} {
	size := len(s.data)
	item := s.data[size-1]
	s.data = s.data[:size-1]
	return item
}

//Peek 获取栈顶元素
func (s *Stack) Peek() interface{} {
	return s.data[len(s.data)-1]
}
