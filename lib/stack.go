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

//IsNotEmpty 判断栈是否不为空
func (s *Stack) IsNotEmpty() bool {
	return len(s.data) != 0
}

//Size 长度
func (s *Stack) Size() int {
	return len(s.data)
}

//Push 新增元素
func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

//Pop 移除栈顶元素
func (s *Stack) Pop() interface{} {
	item := s.data[s.Size()-1]
	s.data = s.data[:s.Size()-1]
	return item
}

//Peek 获取栈顶元素
func (s *Stack) Peek() interface{} {
	return s.data[s.Size()-1]
}
