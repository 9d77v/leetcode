package stack

//Stack 栈
type Stack interface {
	Len() int
	IsEmpty() bool
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	Clear()
}
