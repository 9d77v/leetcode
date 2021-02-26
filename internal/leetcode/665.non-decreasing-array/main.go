package main

/*
题目：非递减数列
给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。

我们是这样定义一个非递减数列的： 对于数组中所有的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-decreasing-array
*/

/*
方法一：贪心（变动数组）
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：36 ms	内存消耗：6.4 MB
*/
func checkPossibility(nums []int) bool {
	count := 0
	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			count++
			if count > 1 {
				return false
			}
			if i > 0 && nums[i-1] > nums[i+1] {
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}
