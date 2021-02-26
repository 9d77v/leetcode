package main

/*
题目：删掉一个元素以后全为1的最长子数组
给你一个二进制数组 nums ，你需要从中删掉一个元素。

请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。

如果不存在这样的子数组，请返回 0 。

提示：

1 <= nums.length <= 10^5
nums[i] 要么是 0 要么是 1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-subarray-of-1s-after-deleting-one-element
*/

/*
方法一：双指针
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：48 ms	内存消耗：8.1 MB
*/
func longestSubarray(nums []int) int {
	left, zeroCnt := 0, 0
	for _, num := range nums {
		if num == 0 {
			zeroCnt++
		}
		if zeroCnt > 1 {
			if nums[left] == 0 {
				zeroCnt--
			}
			left++
		}
	}
	return len(nums) - left - 1
}
