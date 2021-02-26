package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/heap"
)

/*
题目：数据流中的第 K 大元素

设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。

请实现 KthLargest 类：

KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。



提示：
1 <= k <= 104
0 <= nums.length <= 104
-104 <= nums[i] <= 104
-104 <= val <= 104
最多调用 add 方法 104 次
题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-largest-element-in-a-stream
*/

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
