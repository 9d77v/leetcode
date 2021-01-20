package main

import (
	"container/heap"
	"math"
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
方法一：优先队列
时间复杂度：О(m)
空间复杂度：О(1)
运行时间：40 ms	内存消耗：6.5 MB
*/
func maximumProduct(nums []int) int {
	max1 := newMinHeap(3)
	min1 := newMaxHeap(3)
	max2 := newMaxHeap(2)
	min2 := newMinHeap(3)
	hasZero := false
	for _, num := range nums {
		if num > 0 {
			if max1.Len() < 3 {
				max1.push(num)
				max2.push(num)
			} else if num > max1.Peek().(int) {
				max1.pop()
				max1.push(num)
			} else if num < max2.Peek().(int) {
				max2.pop()
				max2.push(num)
			}
		} else if num < 0 {
			if min1.Len() < 2 {
				min1.push(num)
			} else if num < min1.Peek().(int) {
				min1.pop()
				min1.push(num)
			}
			if min2.Len() < 3 {
				min2.push(num)
			} else if num > min2.Peek().(int) {
				min2.pop()
				min2.push(num)
			}
		} else {
			hasZero = true
		}
	}
	if max1.Len() == 3 {
		num1, num2, num3 := max1.pop(), max1.pop(), max1.pop()
		max := num1 * num2 * num3
		if min1.Len() == 2 {
			max = Max(max, num3*min1.pop()*min1.pop())
		}
		return max
	}
	if max1.Len() == 2 {
		if min1.Len() == 2 {
			max1.pop()
			return max1.pop() * min1.pop() * min1.pop()
		}
		if hasZero {
			return 0
		}
		return max1.pop() * max1.pop() * min2.pop()
	}
	if max1.Len() == 1 {
		if min1.Len() == 2 {
			return max1.pop() * min1.pop() * min1.pop()
		}
		return 0
	}
	if hasZero {
		return 0
	}
	return min2.pop() * min2.pop() * min2.pop()
}

type MaxHeap struct {
	*Heap
}

func newMaxHeap(n int) *MaxHeap {
	return &MaxHeap{
		NewHeap(n),
	}
}

func (hp *MaxHeap) Less(i, j int) bool {
	return hp.Data[i].(int) > hp.Data[j].(int)
}

func (hp *MaxHeap) push(i int) {
	heap.Push(hp, i)
}

func (hp *MaxHeap) pop() int {
	return heap.Pop(hp).(int)
}

type MinHeap struct {
	*Heap
}

func newMinHeap(n int) *MinHeap {
	return &MinHeap{
		NewHeap(n),
	}
}

func (hp *MinHeap) Less(i, j int) bool {
	return hp.Data[i].(int) < hp.Data[j].(int)
}

func (hp *MinHeap) push(i int) {
	heap.Push(hp, i)
}

func (hp *MinHeap) pop() int {
	return heap.Pop(hp).(int)
}

/*
方法二：线性扫描
时间复杂度：О(m)
空间复杂度：О(1)
运行时间：36 ms	内存消耗：6.5 MB
*/
func maximumProductFunc2(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MaxInt64
	min1, min2 := math.MaxInt64, math.MaxInt64
	for _, num := range nums {
		if num < min1 {
			min2, min1 = min1, num
		} else if num < min2 {
			min2 = num
		}

		if num > max1 {
			max3, max2, max1 = max2, max1, num
		} else if num > max2 {
			max3, max2 = max2, num
		} else if num > max3 {
			max3 = num
		}
	}
	return Max(min1*min2*max1, max1*max2*max3)
}

/*
方法三：排序
时间复杂度：О(NlogN)
空间复杂度：О(logN)
运行时间：40 ms	内存消耗：6.5 MB
*/
func maximumProductFunc3(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return Max(nums[0]*nums[1]*nums[n-1], nums[n-1]*nums[n-2]*nums[n-3])
}
