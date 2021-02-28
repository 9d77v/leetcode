package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

提示：

1 <= nums.length <= 3 * 104
-105 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof
*/

/*
方法一: 动态规划
时间复杂度：插入О(n)
空间复杂度：О(1)
运行时间：28 ms	内存消耗：7.1 MB
*/
func maxSubArray(nums []int) int {
	n := len(nums)
	prev, cur, max := nums[0], 0, nums[0]
	for i := 1; i < n; i++ {
		if prev < 0 {
			cur = nums[i]
		} else {
			cur = prev + nums[i]
		}
		max = Max(max, cur)
		prev = cur
	}
	return max
}
