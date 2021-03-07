package main

/*
题目：
给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-element-ii
*/

/*
方法一：单调栈+拷贝数组
时间复杂度：О(m+n)
空间复杂度：О(n)
运行时间：32 ms	内存消耗：6.7 MB
*/
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	max, maxIndex := nums[0], 0
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max, maxIndex = nums[i], i
		}
	}
	nums = append(nums, nums[:maxIndex+1]...)
	stack := []int{}
	result := make([]int, n)
	for i, num := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			peek := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if peek < n {
				result[peek] = num
			}
		}
		if num == max && i < n {
			result[i] = -1
		} else {
			stack = append(stack, i)
		}
	}
	return result
}

/*
方法一：单调栈+取模
时间复杂度：О(m+n)
空间复杂度：О(n)
运行时间：36 ms	内存消耗：6.6 MB
*/
func nextGreaterElementsFunc2(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	stack := []int{}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}
	for i := 0; i < 2*n-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			result[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return result
}
