package heap

//DualHeap 大小堆
type DualHeap struct {
	small     Heap
	large     Heap
	delayed   map[int]int
	k         int
	smallSize int
	largeSize int
}

//NewDualHeap 初始化大小堆
func NewDualHeap(k int) *DualHeap {
	n := k>>1 + 2
	return &DualHeap{
		small:   NewMaxHeap(n),
		large:   NewMinHeap(n),
		delayed: make(map[int]int, 0),
		k:       k,
	}
}

//Add 新增元素
func (h *DualHeap) Add(num int) {
	if h.small.IsEmpty() || num < h.small.Peek().(int) {
		h.small.PushItem(num)
		h.smallSize++
	} else {
		h.large.PushItem(num)
		h.largeSize++
	}
	h.makeBalance()
}

//Remove 删除元素
func (h *DualHeap) Remove(num int) {
	h.delayed[num]++
	if num <= h.small.Peek().(int) {
		h.smallSize--
		if num == h.small.Peek().(int) {
			h.prune(h.small)
		}
	} else {
		h.largeSize--
		if num == h.large.Peek().(int) {
			h.prune(h.large)
		}
	}
	h.makeBalance()
}

//Median 获取中位数
func (h *DualHeap) Median() float64 {
	if h.k&1 == 1 {
		return float64(h.small.Peek().(int))
	}
	return (float64(h.small.Peek().(int)) + float64(h.large.Peek().(int))) / 2
}

//makeBalance 保持数量平衡
func (h *DualHeap) makeBalance() {
	if h.smallSize > h.largeSize+1 {
		h.large.PushItem(h.small.PopItem())
		h.smallSize--
		h.largeSize++
		h.prune(h.small)
	} else if h.smallSize < h.largeSize {
		h.small.PushItem(h.large.PopItem())
		h.smallSize++
		h.largeSize--
		h.prune(h.large)
	}
}

//prune 弹出需要删除元素
func (h *DualHeap) prune(hp Heap) {
	for !hp.IsEmpty() {
		num := hp.Peek().(int)
		if _, ok := h.delayed[num]; ok {
			h.delayed[num]--
			if h.delayed[num] == 0 {
				delete(h.delayed, num)
			}
			hp.PopItem()
		} else {
			break
		}
	}
}
