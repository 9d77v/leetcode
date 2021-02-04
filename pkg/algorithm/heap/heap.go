package heap

//Heap 堆
type Heap interface {
	PushItem(v interface{})
	PopItem() interface{}
	Len() int
	Peek() interface{}
	IsEmpty() bool
}
