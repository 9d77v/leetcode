package main

import (
	"math"
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm"
	. "github.com/9d77v/leetcode/pkg/algorithm/heap"
)

/*
方法一：优先队列
时间复杂度：О(m)
空间复杂度：О(1)
运行时间：40 ms	内存消耗：6.5 MB
*/
func maximumProduct(nums []int) int {
	max1 := NewMinHeap(3)
	min1 := NewMaxHeap(3)
	max2 := NewMaxHeap(2)
	min2 := NewMinHeap(3)
	hasZero := false
	for _, num := range nums {
		if num > 0 {
			if max1.Len() < 3 {
				max1.PushItem(num)
				max2.PushItem(num)
			} else if num > max1.Peek().(int) {
				max1.PopItem()
				max1.PushItem(num)
			} else if num < max2.Peek().(int) {
				max2.PopItem()
				max2.PushItem(num)
			}
		} else if num < 0 {
			if min1.Len() < 2 {
				min1.PushItem(num)
			} else if num < min1.Peek().(int) {
				min1.PopItem()
				min1.PushItem(num)
			}
			if min2.Len() < 3 {
				min2.PushItem(num)
			} else if num > min2.Peek().(int) {
				min2.PopItem()
				min2.PushItem(num)
			}
		} else {
			hasZero = true
		}
	}
	if max1.Len() == 3 {
		num1, num2, num3 := max1.PopItem().(int), max1.PopItem().(int), max1.PopItem().(int)
		max := num1 * num2 * num3
		if min1.Len() == 2 {
			max = Max(max, num3*min1.PopItem().(int)*min1.PopItem().(int))
		}
		return max
	}
	if max1.Len() == 2 {
		if min1.Len() == 2 {
			max1.PopItem()
			return max1.PopItem().(int) * min1.PopItem().(int) * min1.PopItem().(int)
		}
		if hasZero {
			return 0
		}
		return max1.PopItem().(int) * max1.PopItem().(int) * min2.PopItem().(int)
	}
	if max1.Len() == 1 {
		if min1.Len() == 2 {
			return max1.PopItem().(int) * min1.PopItem().(int) * min1.PopItem().(int)
		}
		return 0
	}
	if hasZero {
		return 0
	}
	return min2.PopItem().(int) * min2.PopItem().(int) * min2.PopItem().(int)
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
