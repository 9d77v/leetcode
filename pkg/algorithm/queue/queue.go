package queue

//Queue 队列
type Queue interface {
	Len() int
	Empty() bool
	Push(v interface{})
	Pop() interface{}
	BFS(func(v interface{}))
}
