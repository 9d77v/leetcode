package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
提示：

0 <= nums.length <= 100
0 <= nums[i] <= 400
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
*/

/*
方法一：动态规划 f(0)=nums[0] f(1)=max(nums[0],nums[1]) f(i)=max(f(i-2)+nums[i],f(i-1))
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：0 ms	内存消耗：2 MB
*/
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return Max(nums[0], nums[1])
	}
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], Max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = Max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

/*
方法一：动态规划 优化空间 f(0)=nums[0] f(1)=max(nums[0],nums[1]) f(i)=max(f(i-2)+nums[i],f(i-1))
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2 MB
*/
func robFunc2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return Max(nums[0], nums[1])
	}
	p, q, r := nums[0], Max(nums[0], nums[1]), 0
	for i := 2; i < n; i++ {
		r = Max(p+nums[i], q)
		p, q = q, r
	}
	return r
}
