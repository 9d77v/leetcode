package main

/*
题目：
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof
*/

/*
方法一：二分
时间复杂度：О(logn)
空间复杂度：О(1)
运行时间：12 ms	内存消耗：6.1 MB
*/
func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := int(uint(l+r) >> 1)
		if mid != nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
