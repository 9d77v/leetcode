package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/heap"
)

type KthLargest struct {
	hp Heap
	k  int
}

func Constructor(k int, nums []int) KthLargest {
	hp := NewMinHeap(k)
	for _, num := range nums {
		if hp.Len() < k {
			hp.PushItem(num)
		} else if num > hp.Peek().(int) {
			hp.PushItem(num)
			hp.PopItem()
		}
	}
	return KthLargest{
		hp: hp,
		k:  k,
	}
}

func (this *KthLargest) Add(val int) int {
	if this.hp.Len() < this.k {
		this.hp.PushItem(val)
		if this.hp.Len() < this.k {
			return 0
		}
		if this.hp.Len() == this.k {
			return this.hp.Peek().(int)
		}
	}

	top := this.hp.Peek().(int)
	if val > top {
		this.hp.PushItem(val)
		this.hp.PopItem()
		top = this.hp.Peek().(int)
	}
	return top
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
