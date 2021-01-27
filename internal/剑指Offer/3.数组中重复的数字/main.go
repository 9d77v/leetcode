package main

import "sort"

/*
题目：在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
限制：2 <= n <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
*/

/*
方法一：排序后相邻数比较,改变原数组
时间复杂度：О(nlogn)
空间复杂度：О(1)
运行时间：52 ms	内存消耗：8.7 MB
*/
func findRepeatNumberFunc1(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

/*
方法二：原地置换,改变原数组
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：36 ms	内存消耗：8.7 MB
*/
func findRepeatNumberFunc2(nums []int) int {
	for i := range nums {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
