package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

提示：
n == height.length
0 <= n <= 3 * 104
0 <= height[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water/
*/

/*
方法一：单调栈
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：4 ms	内存消耗：3.1 MB
*/
func trapFunc1(height []int) (max int) {
	stack := make([]int, 0, len(height))
	for i, h := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < h {
			topValue := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				d := i - top - 1
				min := Min(height[top], height[i])
				max += d * (min - topValue)
			}
		}
		stack = append(stack, i)
	}
	return
}
