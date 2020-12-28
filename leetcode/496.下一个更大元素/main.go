package main

import (
	. "github.com/9d77v/leetcode/lib"
)

/*
题目：
给定两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

提示：
nums1和nums2中所有元素是唯一的。
nums1和nums2 的数组大小都不超过1000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-element-i
*/

/*
方法一：单调栈
时间复杂度：О(m+n)
空间复杂度：О(n)
运行时间：4 ms	内存消耗：3.4 MB
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = -1
	}
	nums1Map := make(map[int]int)
	for i, v := range nums1 {
		nums1Map[v] = i + 1
	}
	stack := NewMonotonyDecreasingStack(n)
	stack.Execute(nums2, func(topIndex, topValue, i int) {
		if nums1Map[topValue] != 0 {
			arr[nums1Map[topValue]-1] = nums2[i]
		}
	})
	return arr
}
