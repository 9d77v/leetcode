package main

/*
题目：
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
/*
方法一：单调队列
时间复杂度：О(n)
空间复杂度：О(k)
运行时间：16 ms	内存消耗：6.4 MB
*/
func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k+1)
	deque := make([]int, 0, k)
	for i, num := range nums {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < num {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if deque[0] == i-k {
			deque = deque[1:]
		}
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
