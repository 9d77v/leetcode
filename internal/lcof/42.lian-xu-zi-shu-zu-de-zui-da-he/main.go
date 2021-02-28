package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。

提示：

1 <= arr.length <= 10^5
-100 <= arr[i] <= 100

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
