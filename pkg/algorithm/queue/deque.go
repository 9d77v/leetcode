package queue

//Deque 双端队列
type Deque interface {
	PushFront(v interface{})
	PopFront() interface{}
	PushBack(v interface{})
	PopBack() interface{}
	Front() interface{}
	Back() interface{}
	Empty() bool
	Cap() int
	Len() int
}
