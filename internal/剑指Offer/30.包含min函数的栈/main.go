package main

/*
题目：
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
*/

/*
方法一：使用等长辅助栈
运行时间：16 ms	内存消耗：8.5 MB
*/

type MinStack struct {
	data     []int
	sub_data []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:     make([]int, 0),
		sub_data: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.sub_data) == 0 {
		this.sub_data = append(this.sub_data, x)
	} else {
		if x < this.sub_data[len(this.sub_data)-1] {
			this.sub_data = append(this.sub_data, x)
		} else {
			this.sub_data = append(this.sub_data, this.sub_data[len(this.sub_data)-1])
		}
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.sub_data = this.sub_data[:len(this.sub_data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) Min() int {
	return this.sub_data[len(this.sub_data)-1]
}

/*
方法二：只存储最小值的辅助栈
运行时间：20 ms	内存消耗：8.5 MB
*/

type MinStack2 struct {
	data     []int
	sub_data []int
}

/** initialize your data structure here. */
func Constructor2() MinStack {
	return MinStack{
		data:     make([]int, 0),
		sub_data: make([]int, 0),
	}
}

func (this *MinStack2) Push(x int) {
	this.data = append(this.data, x)
	if len(this.sub_data) == 0 || x <= this.Min() {
		this.sub_data = append(this.sub_data, x)
	}
}

func (this *MinStack2) Pop() {
	top := this.Top()
	this.data = this.data[:len(this.data)-1]
	if top == this.Min() {
		this.sub_data = this.sub_data[:len(this.sub_data)-1]
	}
}

func (this *MinStack2) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack2) Min() int {
	return this.sub_data[len(this.sub_data)-1]
}
