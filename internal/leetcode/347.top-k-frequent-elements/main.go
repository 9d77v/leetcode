package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm/heap"
)

/*
题目：前k个高频元素
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

提示：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
你可以按任意顺序返回答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/top-k-frequent-elements
*/

/*
方法一：小顶堆
时间复杂度：О(n㏒k)
空间复杂度：О(n)
运行时间：12 ms	内存消耗：5.3 MB
*/
func topKFrequentFunc1(nums []int, k int) []int {
	freqMap := make(map[int]int, 0)
	for _, v := range nums {
		freqMap[v]++
	}
	var hp Heap = NewMinHeap(k)
	for key, value := range freqMap {
		hp.PushItem([2]int{key, value})
		if hp.Len() > k {
			hp.PopItem()
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = hp.PopItem().([2]int)[0]
	}
	return res
}

/*
方法二：快排变形
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：12 ms	内存消耗：5.3 MB
*/
func topKFrequentFunc2(nums []int, k int) []int {
	freqMap := make(map[int]int, 0)
	for _, v := range nums {
		freqMap[v]++
	}
	values := make([][2]int, 0, len(freqMap))
	for key, value := range freqMap {
		values = append(values, [2]int{key, value})
	}
	res := make([]int, k)
	qsort(values, 0, len(values)-1, res, 0, k)
	return res
}

func qsort(values [][2]int, start, end int, ret []int, retIndex, k int) {
	pivot := values[start][1]
	left := start
	for i := start + 1; i <= end; i++ {
		if values[i][1] >= pivot {
			values[left+1], values[i] = values[i], values[left+1]
			left++
		}
	}
	if k <= left-start {
		qsort(values, start, left-1, ret, retIndex, k)
	} else {
		for i := start; i <= left; i++ {
			ret[retIndex] = values[i][0]
			retIndex++
		}
		if k > left-start+1 {
			qsort(values, left+1, end, ret, retIndex, k-(left-start+1))
		}
	}
}
