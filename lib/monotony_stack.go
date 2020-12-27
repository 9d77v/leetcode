package lib

//MonotonyIncreasingStack 单调递增栈
type MonotonyIncreasingStack struct {
	*Stack
}

//NewMonotonyIncreasingStack 初始化单调递增栈
func NewMonotonyIncreasingStack(len int) *MonotonyIncreasingStack {
	return &MonotonyIncreasingStack{
		Stack: NewStack(len + 2),
	}
}

//Execute 执行单调递增栈
//top: 出栈值
//i: 当前索引
func (m *MonotonyIncreasingStack) Execute(data []int,
	callback func(top, index int)) {
	data = append([]int{0}, append(data, 0)...)
	for i := 0; i < len(data); i++ {
		for m.IsNotEmpty() && data[m.Peek().(int)] > data[i] {
			top := m.Pop().(int)
			if callback != nil {
				callback(data[top], i)
			}
		}
		m.Push(i)
	}
}

//MonotonyDecreasingStack 单调递减栈
type MonotonyDecreasingStack struct {
	*Stack
}

//NewMonotonyDecreasingStack 初始化单调递减栈
func NewMonotonyDecreasingStack() *MonotonyDecreasingStack {
	return &MonotonyDecreasingStack{
		Stack: NewStack(0),
	}
}

//Execute 执行单调递减栈
//top: 出栈值
//i: 当前索引
func (m *MonotonyDecreasingStack) Execute(data []int,
	callback func(top, index int)) {
	data = append([]int{0}, append(data, 0)...)
	for i := len(data) - 1; i >= 0; i-- {
		for m.IsNotEmpty() && data[m.Peek().(int)] < data[i] {
			top := m.Pop().(int)
			if callback != nil {
				callback(data[top], i)
			}
		}
		m.Push(i)
	}
}
