package stack

//SliceStack 切片栈
type SliceStack struct {
	data []interface{}
}

//NewSliceStack 初始化栈
func NewSliceStack(len int) *SliceStack {
	return &SliceStack{
		data: make([]interface{}, 0, len),
	}
}

//Len 长度
func (s *SliceStack) Len() int {
	return len(s.data)
}

//Empty 判断栈是否为空
func (s *SliceStack) Empty() bool {
	return s.Len() == 0
}

//Push 新增元素
func (s *SliceStack) Push(v interface{}) {
	s.data = append(s.data, v)
}

//Pop 移除栈顶元素
func (s *SliceStack) Pop() interface{} {
	if s.Len() == 0 {
		return nil
	}
	item := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]
	return item
}

//Peek 获取栈顶元素
func (s *SliceStack) Peek() interface{} {
	if s.Len() == 0 {
		return nil
	}
	return s.data[s.Len()-1]
}

//Clear 清空栈
func (s *SliceStack) Clear() {
	s.data = s.data[0:0]
}
