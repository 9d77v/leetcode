package main

/*
题目：每日温度
请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/daily-temperatures
*/

/*
方法一：单调栈
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：84 ms	内存消耗：8.8 MB
*/
func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))
	stack := make([]int, 0, len(T))
	for i, t := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < t {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[topIndex] = i - topIndex
		}
		stack = append(stack, i)
	}
	return result
}
