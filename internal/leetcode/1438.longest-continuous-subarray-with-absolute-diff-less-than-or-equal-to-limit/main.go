package main

import (
	"math"

	. "github.com/9d77v/leetcode/pkg/algorithm"
	. "github.com/9d77v/leetcode/pkg/algorithm/treap"
)

/*
题目：绝对差值不超过限制的最长连续子数组
给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。

如果不存在满足条件的子数组，则返回 0 。

提示：

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
0 <= limit <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit
*/

/*
方法一：滑动窗口+map
时间复杂度：О(n^2)
空间复杂度：О(n)
运行时间：1048 ms	内存消耗：7.9 MB
*/
func longestSubarray(nums []int, limit int) int {
	left, min, max, numMap := 0, 0, math.MinInt64, make(map[int]int, 0)
	for _, num := range nums {
		numMap[num]++
		min, max = Min(min, num), Max(max, num)
		if max-min > limit {
			numMap[nums[left]]--
			if numMap[nums[left]] == 0 {
				delete(numMap, nums[left])
			}
			left++
			min, max = getMinMax(numMap)
		}
	}
	return len(nums) - left
}

func getMinMax(numMap map[int]int) (int, int) {
	min, max := math.MaxInt64, math.MinInt64
	for num := range numMap {
		min, max = Min(min, num), Max(max, num)
	}
	return min, max
}

/*
方法一：滑动窗口+treap
时间复杂度：О(n)
空间复杂度：О(n)
运行时间：120 ms	内存消耗：8.7 MB
*/
func longestSubarrayFunc2(nums []int, limit int) int {
	left, treap := 0, new(Treap)
	for _, num := range nums {
		treap.Insert(num)
		if treap.Max()-treap.Min() > limit {
			treap.Delete(nums[left])
			left++
		}
	}
	return len(nums) - left
}
