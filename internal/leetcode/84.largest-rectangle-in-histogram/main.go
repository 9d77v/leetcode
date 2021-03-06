package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
	. "github.com/9d77v/leetcode/pkg/algorithm/stack"
)

/*
题目：柱状图中最大的矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
*/

/*
方法一：暴力方法
时间复杂度：О(n²)
空间复杂度：О(1)
运行时间：688 ms	内存消耗：4.2 MB
*/
func largestRectangleAreaFunc1(heights []int) (max int) {
	n := len(heights)
	for i, height := range heights {
		j, left, right := 1, i, i
		for left-1 > -1 && heights[left-1] >= heights[i] {
			j++
			left--
		}
		for right+1 < n && heights[right+1] >= heights[i] {
			j++
			right++
		}
		max = Max(max, height*j)
	}
	return max
}

/*
方法二：单调栈
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：12 ms	内存消耗：5.7 MB
*/
func largestRectangleAreaFunc2(heights []int) (max int) {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	//栈存放下标
	var stack Stack = NewSliceStack(n)
	for i := 0; i < n; i++ {
		for !stack.IsEmpty() && heights[stack.Peek().(int)] > heights[i] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			left[i] = -1
		} else {
			left[i] = stack.Peek().(int)
		}
		stack.Push(i)
	}
	stack = NewSliceStack(n)
	for i := n - 1; i > -1; i-- {
		for !stack.IsEmpty() && heights[stack.Peek().(int)] > heights[i] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			right[i] = n
		} else {
			right[i] = stack.Peek().(int)
		}
		stack.Push(i)
	}
	for i := 0; i < n; i++ {
		max = Max(max, (right[i]-left[i]-1)*heights[i])
	}
	return
}

/*
方法三：单调栈+常数优化
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：8 ms	内存消耗：5.5 MB
*/
func largestRectangleAreaFunc3(heights []int) (max int) {
	heights = append(heights, -1)
	stack := make([]int, 0, len(heights))
	for i, height := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > height {
			topValue := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			right := i - 1
			left := 0
			if len(stack) > 0 {
				left = stack[len(stack)-1] + 1
			}
			max = Max(max, (right-left+1)*topValue)
		}
		stack = append(stack, i)
	}
	return
}
