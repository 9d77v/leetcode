package main

import "sort"

/*
题目：在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
限制：2 <= n <= 100000
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
*/

/*
方法一：排序后相邻数比较
时间复杂度：О(n)
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
方法二：哈希表
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：40 ms	内存消耗：8 MB
*/
func findRepeatNumberFunc2(nums []int) int {
	arr := make([]int, len(nums))
	for _, num := range nums {
		if arr[num] == 1 {
			return num
		}
		arr[num] = 1
	}
	return -1
}

/*
方法三：原地置换
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：36 ms	内存消耗：8.7 MB
*/
func findRepeatNumberFunc3(nums []int) int {
	for i, num := range nums {
		if i == num {
			continue
		}
		if num == nums[num] {
			return num
		}
		nums[i], nums[num] = nums[num], nums[i]
	}
	return -1
}
