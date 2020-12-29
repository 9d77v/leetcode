package main

import "math"

/*
题目：
假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] 。

请找出其中最小的元素。

提示：
1 <= nums.length <= 5000
-5000 <= nums[i] <= 5000
nums 中的所有整数都是 唯一 的
nums 原来是一个升序排序的数组，但在预先未知的某个点上进行了旋转

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
*/

/*
方法一：二分查找
时间复杂度：О(㏒n)
空间复杂度：О(1)
运行时间：4 ms	内存消耗：3.1 MB
*/
func findMin(nums []int) int {
	n := len(nums)
	if len(nums) == 0 {
		return math.MinInt64
	}
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
