package heap

import (
	"container/heap"
	"log"
)

//BaseHeap 堆
type BaseHeap struct {
	data  []interface{}
	isMax bool //是否大根堆
}

//NewMaxHeap 初始化大顶堆
func NewMaxHeap(n ...int) *BaseHeap {
	cap := 0
	if len(n) > 0 {
		cap = n[0]
	}
	hp := &BaseHeap{
		data:  make([]interface{}, 0, cap),
		isMax: true,
	}
	heap.Init(hp)
	return hp
}

//NewMinHeap 初始化堆
func NewMinHeap(n ...int) *BaseHeap {
	cap := 0
	if len(n) > 0 {
		cap = n[0]
	}
	hp := &BaseHeap{
		data:  make([]interface{}, 0, cap),
		isMax: false,
	}
	heap.Init(hp)
	return hp
}

//Len 长度
func (hp *BaseHeap) Len() int {
	return len(hp.data)
}

//IsEmpty 是否为空
func (hp *BaseHeap) IsEmpty() bool {
	return len(hp.data) == 0
}

//Swap 交换元素
func (hp *BaseHeap) Swap(i, j int) {
	hp.data[i], hp.data[j] = hp.data[j], hp.data[i]
}

//Less 比较大小
func (hp *BaseHeap) Less(i, j int) bool {
	switch hp.data[i].(type) {
	case int:
		if hp.isMax {
			return hp.data[i].(int) > hp.data[j].(int)
		}
		return hp.data[i].(int) < hp.data[j].(int)
	case byte:
		if hp.isMax {
			return hp.data[i].(byte) > hp.data[j].(byte)
		}
		return hp.data[i].(byte) < hp.data[j].(byte)
	case [2]int:
		if hp.isMax {
			return hp.data[i].([2]int)[1] > hp.data[j].([2]int)[1]
		}
		return hp.data[i].([2]int)[1] < hp.data[j].([2]int)[1]
	default:
		log.Panic("not supported type")
	}
	return false
}

//Push 新增元素
func (hp *BaseHeap) Push(v interface{}) {
	hp.data = append(hp.data, v)
}

//Pop 移除堆顶元素
func (hp *BaseHeap) Pop() interface{} {
	item := hp.data[hp.Len()-1]
	hp.data = hp.data[:hp.Len()-1]
	return item
}

//PushItem 新增元素
func (hp *BaseHeap) PushItem(v interface{}) {
	heap.Push(hp, v)
}

//PopItem 移除堆顶元素
func (hp *BaseHeap) PopItem() interface{} {
	return heap.Pop(hp)
}

//Peek 获取堆顶元素
func (hp *BaseHeap) Peek() interface{} {
	return hp.data[0]
}
