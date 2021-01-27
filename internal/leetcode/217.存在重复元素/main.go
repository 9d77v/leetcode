package main

import "sort"

/*
题目：
给定一个整数数组，判断是否存在重复元素。

如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
*/

/*
方法一：哈希表
时间复杂度：O(n)
空间复杂度：O(n)
运行时间：20 ms	内存消耗：7.9 MB
*/
func containsDuplicateFunc1(nums []int) bool {
	dataMap := make(map[int]struct{}, 0)
	for _, num := range nums {
		if _, ok := dataMap[num]; !ok {
			dataMap[num] = struct{}{}
		} else {
			return true
		}
	}
	return false
}

/*
方法二：排序
时间复杂度：O(nlogn)
空间复杂度：O(logn)
运行时间：20 ms	内存消耗：6 MB
*/
func containsDuplicateFunc2(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
