package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
给定一个二进制数组， 计算其中最大连续1的个数。
*/

/*
方法一：双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2.1 MB
*/
func findMaxConsecutiveOnes(nums []int) int {
	left, zeroCount := 0, 0
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		}
		if zeroCount > 0 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
	}
	return len(nums) - left
}

/*
方法二：一次遍历
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：0 ms	内存消耗：2.1 MB
*/
func findMaxConsecutiveOnesFunc2(nums []int) int {
	maxCount, count := 0, 0
	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			maxCount = Max(maxCount, count)
			count = 0
		}
	}
	maxCount = Max(maxCount, count)
	return maxCount
}
