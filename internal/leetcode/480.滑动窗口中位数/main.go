package main

import (
	"container/heap"
	"sort"

	. "github.com/9d77v/leetcode/pkg/algorithm"
	. "github.com/9d77v/leetcode/pkg/algorithm/heap"
)

/*
题目：
中位数是有序序列最中间的那个数。如果序列的长度是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。

例如：

[2,3,4]，中位数是 3
[2,3]，中位数是 (2 + 3) / 2 = 2.5
给你一个数组 nums，有一个长度为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。

提示：

你可以假设 k 始终有效，即：k 始终小于输入的非空数组的元素个数。
与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sliding-window-median
*/

/*
方法一：双指针
时间复杂度：О(nklogk)
空间复杂度：О(k)
运行时间：超出时间限制 ms	内存消耗： MB
*/
func medianSlidingWindow(nums []int, k int) []float64 {
	left, right, result := 0, k-1, make([]float64, len(nums)-k+1)
	for ; right < len(nums); right++ {
		t := append([]int(nil), nums[left:left+k]...)
		sort.Ints(t)
		result[left] = float64(Median(t))
		left++
	}
	return result
}

/*
方法二：两个优先队列
时间复杂度：О(nk)
空间复杂度：О(n)
运行时间：36ms	内存消耗： 6.6MB
*/
func medianSlidingWindowFunc2(nums []int, k int) []float64 {
	delayed = map[int]int{} // 哈希表，记录「延迟删除」的元素，key 为元素，value 为需要删除的次数
	small = &hp{}           // 大根堆，维护较小的一半元素
	large = &hp{}           // 小根堆，维护较大的一半元素
	makeBalance := func() {
		// 调整 small 和 large 中的元素个数，使得二者的元素个数满足要求
		if small.size > large.size+1 { // small 比 large 元素多 2 个
			large.push(-small.pop())
			small.prune() // small 堆顶元素被移除，需要进行 prune
		} else if small.size < large.size { // large 比 small 元素多 1 个
			small.push(-large.pop())
			large.prune() // large 堆顶元素被移除，需要进行 prune
		}
	}
	insert := func(num int) {
		if small.Len() == 0 || num <= -small.IntSlice[0] {
			small.push(-num)
		} else {
			large.push(num)
		}
		makeBalance()
	}
	erase := func(num int) {
		delayed[num]++
		if num <= -small.IntSlice[0] {
			small.size--
			if num == -small.IntSlice[0] {
				small.prune()
			}
		} else {
			large.size--
			if num == large.IntSlice[0] {
				large.prune()
			}
		}
		makeBalance()
	}
	getMedian := func() float64 {
		if k&1 > 0 {
			return float64(-small.IntSlice[0])
		}
		return float64(-small.IntSlice[0]+large.IntSlice[0]) / 2
	}

	for _, num := range nums[:k] {
		insert(num)
	}
	n := len(nums)
	ans := make([]float64, 1, n-k+1)
	ans[0] = getMedian()
	for i := k; i < n; i++ {
		insert(nums[i])
		erase(nums[i-k])
		ans = append(ans, getMedian())
	}
	return ans
}

type hp struct {
	sort.IntSlice
	size int
}

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) { h.size++; heap.Push(h, v) }
func (h *hp) pop() int   { h.size--; return heap.Pop(h).(int) }
func (h *hp) prune() {
	for h.Len() > 0 {
		num := h.IntSlice[0]
		if h == small {
			num = -num
		}
		if d, has := delayed[num]; has {
			if d > 1 {
				delayed[num]--
			} else {
				delete(delayed, num)
			}
			heap.Pop(h)
		} else {
			break
		}
	}
}

var delayed map[int]int
var small, large *hp

/*
方法三：两个优先队列
时间复杂度：О(nk)
空间复杂度：О(n)
运行时间：32ms	内存消耗： 6.8MB
*/
func medianSlidingWindowFunc3(nums []int, k int) []float64 {
	n := len(nums)
	h := NewDualHeap(k)
	for i := 0; i < k; i++ {
		h.Add(nums[i])
	}
	result := make([]float64, n-k+1)
	result[0] = h.Median()
	for i := k; i < n; i++ {
		h.Add(nums[i])
		h.Remove(nums[i-k])
		result[i-k+1] = h.Median()
	}
	return result
}
