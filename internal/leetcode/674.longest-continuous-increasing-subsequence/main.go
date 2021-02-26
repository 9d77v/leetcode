package main

import (
	"math"

	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：最长连续递增序列
给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。

连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。

提示：

0 <= nums.length <= 104
-109 <= nums[i] <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence
*/

/*
方法一：贪心
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：8 ms	内存消耗：4.2 MB
*/
func findLengthOfLCISFunc1(nums []int) int {
	prev, num := math.MinInt64, 0
	result := 0
	for i := 0; i < len(nums); i++ {
		if prev < nums[i] {
			num++
		}
		if i == len(nums)-1 {
			result = Max(result, num)
		} else if prev >= nums[i] {
			result = Max(result, num)
			num = 1
		}
		prev = nums[i]
	}
	return result
}

/*
方法二：贪心
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：8 ms	内存消耗：4.2 MB
*/
func findLengthOfLCISFunc2(nums []int) int {
	result := 0
	n := len(nums)
	for i := 0; i < n; {
		j := i
		for j+1 < n && nums[j] < nums[j+1] {
			j++
		}
		result = Max(result, j-i+1)
		i = j + 1
	}
	return result
}
