package lib

//MonotonyDecreasingStack 单调递减栈
type MonotonyDecreasingStack struct {
	*Stack
}

//NewMonotonyDecreasingStack 初始化单调递减栈
func NewMonotonyDecreasingStack(len int) *MonotonyDecreasingStack {
	return &MonotonyDecreasingStack{
		Stack: NewStack(len),
	}
}

//Execute 执行单调递减栈
//top: 出栈值
//i: 当前索引
func (m *MonotonyDecreasingStack) Execute(data []int, fn func(topIndex, topValue, i int)) {
	for i := 0; i < len(data); i++ {
		for m.IsNotEmpty() && data[m.Peek().(int)] < data[i] {
			top := m.Pop().(int)
			if fn != nil {
				fn(top, data[top], i)
			}
		}
		m.Push(i)
	}
}
