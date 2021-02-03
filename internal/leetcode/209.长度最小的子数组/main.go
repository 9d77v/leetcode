package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum
*/

/*
方法一：双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：12 ms	内存消耗：3.7 MB
*/
func minSubArrayLenFunc1(s int, nums []int) int {
	left, res, max := 0, 0, 0
	for right, num := range nums {
		max += num
		for max >= s {
			if res == 0 {
				res = right - left + 1
			} else {
				res = Min(res, right-left+1)
			}
			max -= nums[left]
			left++
		}
	}
	return res
}
