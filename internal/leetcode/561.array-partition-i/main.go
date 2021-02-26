package main

import "sort"

/*
题目：数组拆分I
给定长度为 2n 的整数数组 nums ，你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从 1 到 n 的 min(ai, bi) 总和最大。

返回该 最大总和 。

提示：

1 <= n <= 104
nums.length == 2 * n
-104 <= nums[i] <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/array-partition-i
*/

/*
方法一：排序取奇数
时间复杂度：О(nlogn)
空间复杂度：О(logn)
运行时间：36 ms	内存消耗：6.6 MB
*/
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}
