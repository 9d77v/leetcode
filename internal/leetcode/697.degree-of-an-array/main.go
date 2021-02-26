package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：数组的度
给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/degree-of-an-array
*/

/*
方法一：滑动窗口
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：28 ms	内存消耗：6.9 MB
*/
func findShortestSubArray(nums []int) int {
	freqMap := make(map[int]int, 0)
	for _, num := range nums {
		freqMap[num]++
	}
	degree := 0
	for _, v := range freqMap {
		degree = Max(degree, v)
	}
	left, curFreqMap := 0, make(map[int]int, 0)
	result := 0
	for right, num := range nums {
		curFreqMap[num]++
		for curFreqMap[num] == degree {
			if result == 0 {
				result = right - left + 1
			} else {
				result = Min(result, right-left+1)
			}
			curFreqMap[nums[left]]--
			left++
		}
	}
	return result
}
