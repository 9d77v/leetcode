package queue

//MonotonicQueue 单调队列
type MonotonicQueue struct {
	Deque
	IsAsc bool
}

//NewMonotonicQueue 初始化单调队列
func NewMonotonicQueue(deque Deque, isAsc bool) *MonotonicQueue {
	return &MonotonicQueue{
		Deque: deque,
		IsAsc: isAsc,
	}
}

//Execute 添加元素
func (m *MonotonicQueue) Execute(data []int, fn func(fronIndex, frontValue int)) {
	for i, v := range data {
		for !m.Empty() && m.compare(data[m.Back().(int)], v) {
			m.PopBack()
		}
		m.PushBack(i)
		if i >= m.Cap()-1 {
			if fn != nil {
				fn(m.Front().(int), data[m.Front().(int)])
			}
			if !m.Empty() && m.Front().(int) == i-m.Cap()+1 {
				m.PopFront()
			}
		}
	}
}

func (m *MonotonicQueue) compare(i, j int) bool {
	if m.IsAsc {
		return i > j
	}
	return i < j
}
