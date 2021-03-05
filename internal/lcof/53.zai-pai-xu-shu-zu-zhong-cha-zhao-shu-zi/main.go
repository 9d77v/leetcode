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
方法一：二分查找后左右扫描
时间复杂度：О(n)
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

/*
方法二：二分查找左右边界
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：8 ms	内存消耗：3.9 MB
*/
func searchFunc2(nums []int, target int) int {
	return binarySearch(nums, target) - binarySearch(nums, target-1)
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := int(uint(l+r) >> 1)
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
