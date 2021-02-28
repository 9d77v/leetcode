package main

import "container/heap"

/*
题目：
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof
*/

/*
方法一: 最大堆存小的[n/2],最小堆存大的[n/2]
时间复杂度：插入О(logN) 查找О(1)
空间复杂度：О(n)
运行时间：144 ms	内存消耗：13.3 MB
*/

type MedianFinder struct {
	maxHeap Heap
	minHeap Heap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: NewMaxHeap(),
		minHeap: NewMinHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == this.minHeap.Len() {
		this.minHeap.PushItem(num)
		this.maxHeap.PushItem(this.minHeap.PopItem())
	} else {
		this.maxHeap.PushItem(num)
		this.minHeap.PushItem(this.maxHeap.PopItem())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		return float64(this.maxHeap.Peek()+this.minHeap.Peek()) / 2
	}
	return float64(this.maxHeap.Peek())
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type Heap interface {
	Len() int
	PushItem(v int)
	PopItem() int
	Peek() int
}

type BaseHeap struct {
	data  []int
	isMax bool
}

func NewMaxHeap() *BaseHeap {
	hp := &BaseHeap{
		data:  make([]int, 0),
		isMax: true,
	}
	heap.Init(hp)
	return hp
}

func NewMinHeap() *BaseHeap {
	hp := &BaseHeap{
		data:  make([]int, 0),
		isMax: false,
	}
	heap.Init(hp)
	return hp
}

func (hp *BaseHeap) Len() int {
	return len(hp.data)
}

func (hp *BaseHeap) PushItem(v int) {
	heap.Push(hp, v)
}

func (hp *BaseHeap) PopItem() int {
	return heap.Pop(hp).(int)
}

func (hp *BaseHeap) Peek() int {
	return hp.data[0]
}

func (hp *BaseHeap) Less(i, j int) bool {
	if hp.isMax {
		return hp.data[i] > hp.data[j]
	}
	return hp.data[i] < hp.data[j]
}

func (hp *BaseHeap) Swap(i, j int) {
	hp.data[i], hp.data[j] = hp.data[j], hp.data[i]
}

func (hp *BaseHeap) Push(v interface{}) {
	hp.data = append(hp.data, v.(int))
}

func (hp *BaseHeap) Pop() interface{} {
	peek := hp.data[len(hp.data)-1]
	hp.data = hp.data[:len(hp.data)-1]
	return peek
}
