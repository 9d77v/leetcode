package main

import "sort"

/*
题目：在排序数组中查找数字 I
统计一个数字在排序数组中出现的次数。
限制：

0 <= 数组长度 <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
*/

/*
方法一：二分查找
时间复杂度：О(logn)
空间复杂度：О(1)
运行时间：8 ms	内存消耗：3.9 MB
*/
func search(nums []int, target int) int {
	i := sort.SearchInts(nums, target)
	if i == len(nums) {
		return 0
	}
	count := 0
	for j := i; j >= 0; j-- {
		if nums[j] == target {
			count++
		} else {
			break
		}
	}
	for j := i + 1; j < len(nums); j++ {
		if nums[j] == target {
			count++
		} else {
			break
		}
	}
	return count
}
