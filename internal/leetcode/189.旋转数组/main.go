package main

import (
	. "github.com/9d77v/leetcode/pkg/algorithm"
)

/*
题目：
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-array/
*/

/*
方法一：暴力解法
时间复杂度：О(nk)
空间复杂度：О(1)
运行时间：104 ms	内存消耗：3.2 MB
*/
func rotateFunc1(nums []int, k int) {
	k %= len(nums)
	for i := 0; i < k; i++ {
		prev := nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			nums[j], prev = prev, nums[j]
		}
	}
}

/*
方法二：环状替换
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：4 ms	内存消耗：3.1 MB
*/
func rotateFunc2(nums []int, k int) {
	n := len(nums)
	k %= n
	for i := 0; i < Gcd(k, n); i++ {
		pre, cur := nums[i], i
		for {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
			if cur == i {
				break
			}
		}
	}
}

/*
方法三：数组反转
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：4 ms	内存消耗：3.1 MB
*/
func rotateFunc3(nums []int, k int) {
	k %= len(nums)
	Reverse(nums)
	Reverse(nums[:k])
	Reverse(nums[k:])
}
