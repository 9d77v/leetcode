package stack

//MonotonicStack 单调栈
type MonotonicStack struct {
	Stack
	IsAsc bool
}

//NewMonotonicIncreasingStack 初始化单调递增栈
func NewMonotonicIncreasingStack(stack Stack) *MonotonicStack {
	return &MonotonicStack{
		Stack: stack,
		IsAsc: true,
	}
}

//NewMonotonicDecreasingStack 初始化单调递减栈
func NewMonotonicDecreasingStack(stack Stack) *MonotonicStack {
	return &MonotonicStack{
		Stack: stack,
		IsAsc: false,
	}
}

//Execute 存index
func (m *MonotonicStack) Execute(data []int, fn func(topIndex, topValue, i int)) {
	for i := 0; i < len(data); i++ {
		for !m.IsEmpty() && m.compare(data[m.Peek().(int)], data[i]) {
			top := m.Pop().(int)
			if fn != nil {
				fn(top, data[top], i)
			}
		}
		m.Push(i)
	}
}

func (m *MonotonicStack) compare(i, j int) bool {
	if m.IsAsc {
		return i > j
	}
	return i < j
}
