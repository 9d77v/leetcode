package algorithm

//MonotonyIncreasingStack 单调递增栈
type MonotonyIncreasingStack struct {
	*Stack
}

//NewMonotonyIncreasingStack 初始化单调递增栈
func NewMonotonyIncreasingStack(len int) *MonotonyIncreasingStack {
	return &MonotonyIncreasingStack{
		Stack: NewStack(len),
	}
}

//Execute 执行单调递增栈
//top: 出栈值
//i: 当前索引
func (m *MonotonyIncreasingStack) Execute(data []int, fn func(topIndex, topValue, i int)) {
	for i := 0; i < len(data); i++ {
		for m.IsNotEmpty() && data[m.Peek().(int)] > data[i] {
			top := m.Pop().(int)
			if fn != nil {
				fn(top, data[top], i)
			}
		}
		m.Push(i)
	}
}
