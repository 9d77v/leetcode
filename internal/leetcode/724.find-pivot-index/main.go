package main

/*
题目：寻找数组的中心索引
给定一个整数类型的数组 nums，请编写一个能够返回数组 “中心索引” 的方法。

我们是这样定义数组 中心索引 的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。

如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个

说明：

nums 的长度范围为 [0, 10000]。
任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-pivot-index
*/

/*
方法一：前缀和
时间复杂度：О(n)
空间复杂度：О(1)
运行时间：24 ms	内存消耗：6.1 MB
*/
func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	left := 0
	for i, num := range nums {
		if left == sum-left-num {
			return i
		}
		left += num

	}
	return -1
}
