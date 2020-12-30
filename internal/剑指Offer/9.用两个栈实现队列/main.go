package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

提示：
1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
*/

/*
运行时间：236 ms	内存消耗：8.3 MB
*/

type CQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func Constructor() CQueue {
	return CQueue{
		stack1: NewStack(0),
		stack2: NewStack(0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.IsEmpty() {
		if this.stack1.IsEmpty() {
			return -1
		}
		for this.stack1.IsNotEmpty() {
			this.stack2.Push(this.stack1.Pop())
		}
	}
	return this.stack2.Pop().(int)
}
