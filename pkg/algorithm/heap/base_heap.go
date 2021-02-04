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
func (h *BaseHeap) Len() int {
	return len(h.data)
}

//IsEmpty 是否为空
func (h *BaseHeap) IsEmpty() bool {
	return len(h.data) == 0
}

//Swap 交换元素
func (h *BaseHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

//Less 比较大小
func (h *BaseHeap) Less(i, j int) bool {
	switch h.data[i].(type) {
	case int:
		if h.isMax {
			return h.data[i].(int) > h.data[j].(int)
		}
		return h.data[i].(int) < h.data[j].(int)
	case byte:
		if h.isMax {
			return h.data[i].(byte) > h.data[j].(byte)
		}
		return h.data[i].(byte) < h.data[j].(byte)
	case [2]int:
		if h.isMax {
			return h.data[i].([2]int)[1] > h.data[j].([2]int)[1]
		}
		return h.data[i].([2]int)[1] < h.data[j].([2]int)[1]
	default:
		log.Panic("not supported type")
	}
	return false
}

//Push 新增元素
func (h *BaseHeap) Push(v interface{}) {
	h.data = append(h.data, v)
}

//Pop 移除堆顶元素
func (h *BaseHeap) Pop() interface{} {
	if h.Len() == 0 {
		return nil
	}
	item := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return item
}

//PushItem 新增元素
func (h *BaseHeap) PushItem(v interface{}) {
	heap.Push(h, v)
}

//PopItem 移除堆顶元素
func (h *BaseHeap) PopItem() interface{} {
	return heap.Pop(h)
}

//Peek 获取堆顶元素
func (h *BaseHeap) Peek() interface{} {
	return h.data[0]
}
