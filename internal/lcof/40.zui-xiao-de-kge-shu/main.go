package main

import (
	"container/heap"
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：

输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

限制：

0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
*/

/*
方法一: 最大堆
时间复杂度：О(nlogk)
空间复杂度：О(k)
运行时间：36 ms	内存消耗：6.6 MB
*/
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	maxHeap := NewMaxHeap(k)
	for i := 0; i < k; i++ {
		maxHeap.PushItem(arr[i])
	}
	for i := k; i < len(arr); i++ {
		if arr[i] < maxHeap.Peek() {
			maxHeap.PopItem()
			maxHeap.PushItem(arr[i])
		}
	}
	return maxHeap.data
}

type Heap interface {
	PushItem(v int)
	PopItem() int
	Len() int
	Peek() int
	IsEmpty() bool
}

type MaxHeap struct {
	data []int
}

func NewMaxHeap(k int) *MaxHeap {
	hp := &MaxHeap{
		data: make([]int, 0, k),
	}
	heap.Init(hp)
	return hp
}

func (hp *MaxHeap) Push(v interface{}) {
	hp.data = append(hp.data, v.(int))
}

func (hp *MaxHeap) Pop() interface{} {
	peek := hp.Peek()
	hp.data = hp.data[:len(hp.data)-1]
	return peek
}

func (hp *MaxHeap) Len() int {
	return len(hp.data)
}

func (hp *MaxHeap) Peek() int {
	return hp.data[0]
}

func (hp *MaxHeap) IsEmpty() bool {
	return len(hp.data) == 0
}

func (hp *MaxHeap) PushItem(v int) {
	heap.Push(hp, v)
}

func (hp *MaxHeap) PopItem() int {
	return heap.Pop(hp).(int)
}

func (hp *MaxHeap) Less(i, j int) bool {
	return hp.data[i] > hp.data[j]
}

func (hp *MaxHeap) Swap(i, j int) {
	hp.data[i], hp.data[j] = hp.data[j], hp.data[i]
}

/*
方法二: 排序
时间复杂度：О(nlogn)
空间复杂度：О(logn)
运行时间：40 ms	内存消耗：6.6 MB
*/
func getLeastNumbersFunc2(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

/*
方法三: 快排变形
时间复杂度：О(nlogn)
空间复杂度：О(logn)
运行时间：28 ms	内存消耗：6.5 MB
*/
func getLeastNumbersFunc3(arr []int, k int) []int {
	n := len(arr)
	if n == k {
		return arr
	}
	if n < k || k <= 0 || n == 0 {
		return []int{}
	}
	QuickSelect(arr, k, 0, n-1)
	return arr[:k]
}
