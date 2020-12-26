package lib

//Stack 栈
type Stack struct {
	data []interface{}
	top  int
}

//NewStack 初始化栈
func NewStack(len int) *Stack {
	return &Stack{
		data: make([]interface{}, len),
		top:  -1,
	}
}

//IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

//Push 新增元素
func (s *Stack) Push(v interface{}) {
	s.top++
	s.data[s.top] = v
}

//Pop 移除栈顶元素
func (s *Stack) Pop() interface{} {
	item := s.data[s.top]
	s.data[s.top] = nil
	s.top--
	return item
}

//Peek 获取栈顶元素
func (s *Stack) Peek() interface{} {
	return s.data[s.top]
}
