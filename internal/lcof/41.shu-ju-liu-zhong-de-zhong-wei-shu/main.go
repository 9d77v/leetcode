package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/heap"
)

/*
题目：
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。

限制：

最多会对 addNum、findMedian 进行 50000 次调用。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof
*/

/*
方法一: 最大堆存小的[n/2],最小堆存大的[n/2]
时间复杂度：插入О(logN) 查找О(1)
空间复杂度：О(n)
运行时间：112 ms	内存消耗：13.1 MB
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
		this.maxHeap.PushItem(num)
		this.minHeap.PushItem(this.maxHeap.PopItem())
	} else {
		this.minHeap.PushItem(num)
		this.maxHeap.PushItem(this.minHeap.PopItem())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		return float64(this.maxHeap.Peek().(int)+this.minHeap.Peek().(int)) / 2
	}
	return float64(this.maxHeap.Peek().(int))
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
