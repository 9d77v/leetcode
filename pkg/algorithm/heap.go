package algorithm

//Heap 堆
type Heap struct {
	Data []interface{}
}

//NewHeap 初始化堆
func NewHeap(len int) *Heap {
	return &Heap{
		Data: make([]interface{}, 0, len),
	}
}

//Len 长度
func (h *Heap) Len() int {
	return len(h.Data)
}

//Swap 交换元素
func (h *Heap) Swap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

//Push 新增元素
func (h *Heap) Push(v interface{}) {
	h.Data = append(h.Data, v)
}

//Pop 移除堆顶元素
func (h *Heap) Pop() interface{} {
	item := h.Data[h.Len()-1]
	h.Data = h.Data[:h.Len()-1]
	return item
}
